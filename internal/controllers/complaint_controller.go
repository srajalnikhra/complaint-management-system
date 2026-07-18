package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/srajalnikhra/complaint-management-system/internal/dto"
	"github.com/srajalnikhra/complaint-management-system/internal/models"
	"github.com/srajalnikhra/complaint-management-system/internal/services"
	"github.com/srajalnikhra/complaint-management-system/internal/utils"
	"github.com/srajalnikhra/complaint-management-system/internal/validation"
)

var complaintService = services.NewComplaintService()

// CreateComplaint handles submission of a new complaint by a user.
// It validates input fields and links the complaint to the logged-in user.
//
// # CreateComplaint godoc
//
// @Summary Create Complaint
// @Description Create a new complaint
// @Tags Complaints
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateComplaintRequest true "Complaint Details"
// @Success 201 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Failure 401 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /complaints [post]
func CreateComplaint(w http.ResponseWriter, r *http.Request) {

	var req dto.CreateComplaintRequest

	// Decode the request body to get the complaint title and description.
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Validate the input parameters.
	msg := validation.ValidateComplaint(req.Title, req.Description)
	if msg != "" {
		utils.Error(w, http.StatusBadRequest, msg)
		return
	}

	// Retrieve the logged-in user ID from the request context.
	userID := r.Context().Value("userID").(int)

	complaint := models.Complaint{
		UserID:      userID,
		Title:       req.Title,
		Description: req.Description,
	}

	// Save the complaint using the service layer.
	if err := complaintService.Create(&complaint); err != nil {
		log.Printf("Create Complaint Error: %v", err)
		utils.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(
		w,
		http.StatusCreated,
		"Complaint created successfully",
		complaint,
	)
}

// GetMyComplaints lists all complaints created by the logged-in user.
func GetMyComplaints(w http.ResponseWriter, r *http.Request) {

	// Retrieve the user ID from the request context.
	userID := r.Context().Value("userID").(int)

	// Fetch complaints belonging to this user.
	complaints, err := complaintService.GetByUserID(userID)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, "Failed to fetch complaints")
		return
	}

	var response []dto.ComplaintResponse

	// Translate the database models to API response structures.
	for _, complaint := range complaints {
		response = append(response, dto.ComplaintResponse{
			ID:          complaint.ID,
			Title:       complaint.Title,
			Description: complaint.Description,
			Status:      complaint.Status,
			CreatedAt:   complaint.CreatedAt,
		})
	}

	utils.Success(
		w,
		http.StatusOK,
		"Complaints fetched successfully",
		response,
	)
}

// GetComplaintByID returns the details of a single complaint.
// It ensures that only the creator of the complaint can access it.
//
// # GetComplaintByID godoc
//
// @Summary Get Complaint By ID
// @Description Get a complaint by its ID
// @Tags Complaints
// @Produce json
// @Security BearerAuth
// @Param id path int true "Complaint ID"
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Failure 401 {object} utils.APIResponse
// @Failure 403 {object} utils.APIResponse
// @Failure 404 {object} utils.APIResponse
// @Router /complaints/{id} [get]
func GetComplaintByID(w http.ResponseWriter, r *http.Request) {

	// Parse the complaint ID from the URL path.
	idStr := strings.TrimPrefix(r.URL.Path, "/complaints/")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid complaint ID")
		return
	}

	// Retrieve the complaint details.
	complaint, err := complaintService.GetByID(id)
	if err != nil {
		utils.Error(w, http.StatusNotFound, err.Error())
		return
	}

	userID := r.Context().Value("userID").(int)

	// Make sure the logged-in user is the owner of the complaint.
	if complaint.UserID != userID {
		utils.Error(w, http.StatusForbidden, "Forbidden")
		return
	}

	// Translate the database model to an API response.
	response := dto.ComplaintResponse{
		ID:          complaint.ID,
		Title:       complaint.Title,
		Description: complaint.Description,
		Status:      complaint.Status,
		CreatedAt:   complaint.CreatedAt,
	}

	utils.Success(
		w,
		http.StatusOK,
		"Complaint fetched successfully",
		response,
	)
}

// UpdateComplaint modifies the details of an existing complaint.
// It verifies ownership and validates input fields before saving changes.
//
// # UpdateComplaint godoc
//
// @Summary Update Complaint
// @Description Update an existing complaint
// @Tags Complaints
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Complaint ID"
// @Param request body dto.UpdateComplaintRequest true "Updated Complaint"
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Failure 401 {object} utils.APIResponse
// @Failure 403 {object} utils.APIResponse
// @Failure 404 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /complaints/{id} [put]
func UpdateComplaint(w http.ResponseWriter, r *http.Request) {

	// Parse the complaint ID from the URL path.
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/complaints/"))
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid complaint ID")
		return
	}

	userID := r.Context().Value("userID").(int)

	// Fetch the existing database record.
	complaint, err := complaintService.GetByID(id)
	if err != nil {
		utils.Error(w, http.StatusNotFound, err.Error())
		return
	}

	// Make sure the logged-in user is the owner of the complaint.
	if complaint.UserID != userID {
		utils.Error(w, http.StatusForbidden, "Forbidden")
		return
	}

	var req dto.UpdateComplaintRequest

	// Decode the updated fields from the request body.
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Validate the updated parameters.
	msg := validation.ValidateComplaint(req.Title, req.Description)
	if msg != "" {
		utils.Error(w, http.StatusBadRequest, msg)
		return
	}

	complaint.Title = req.Title
	complaint.Description = req.Description

	// Save the changes to the database.
	if err := complaintService.Update(complaint); err != nil {
		utils.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(
		w,
		http.StatusOK,
		"Complaint updated successfully",
		complaint,
	)
}

// DeleteComplaint removes a complaint from the database.
// It ensures only the owner can delete the complaint.
//
// # DeleteComplaint godoc
//
// @Summary Delete Complaint
// @Description Delete a complaint by ID
// @Tags Complaints
// @Produce json
// @Security BearerAuth
// @Param id path int true "Complaint ID"
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Failure 401 {object} utils.APIResponse
// @Failure 403 {object} utils.APIResponse
// @Failure 404 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /complaints/{id} [delete]
func DeleteComplaint(w http.ResponseWriter, r *http.Request) {

	// Parse the complaint ID from the URL path.
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/complaints/"))
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid complaint ID")
		return
	}

	userID := r.Context().Value("userID").(int)

	// Fetch the complaint to verify ownership.
	complaint, err := complaintService.GetByID(id)
	if err != nil {
		utils.Error(w, http.StatusNotFound, err.Error())
		return
	}

	// Make sure the logged-in user is the owner of the complaint.
	if complaint.UserID != userID {
		utils.Error(w, http.StatusForbidden, "Forbidden")
		return
	}

	// Delete the complaint from the database.
	if err := complaintService.Delete(id); err != nil {
		utils.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(
		w,
		http.StatusOK,
		"Complaint deleted successfully",
		nil,
	)
}
