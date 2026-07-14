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

// CreateComplaint godoc
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

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	msg := validation.ValidateComplaint(req.Title, req.Description)
	if msg != "" {
		utils.Error(w, http.StatusBadRequest, msg)
		return
	}

	userID := r.Context().Value("userID").(int)

	complaint := models.Complaint{
		UserID:      userID,
		Title:       req.Title,
		Description: req.Description,
	}

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

// GetMyComplaints godoc
//
// @Summary Get My Complaints
// @Description Get all complaints created by the logged-in user
// @Tags Complaints
// @Produce json
// @Security BearerAuth
// @Success 200 {object} utils.APIResponse
// @Failure 401 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /complaints [get]
func GetMyComplaints(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value("userID").(int)

	complaints, err := complaintService.GetByUserID(userID)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, "Failed to fetch complaints")
		return
	}

	var response []dto.ComplaintResponse

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

// GetComplaintByID godoc
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

	idStr := strings.TrimPrefix(r.URL.Path, "/complaints/")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid complaint ID")
		return
	}

	complaint, err := complaintService.GetByID(id)
	if err != nil {
		utils.Error(w, http.StatusNotFound, err.Error())
		return
	}

	userID := r.Context().Value("userID").(int)

	if complaint.UserID != userID {
		utils.Error(w, http.StatusForbidden, "Forbidden")
		return
	}

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

// UpdateComplaint godoc
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

	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/complaints/"))
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid complaint ID")
		return
	}

	userID := r.Context().Value("userID").(int)

	complaint, err := complaintService.GetByID(id)
	if err != nil {
		utils.Error(w, http.StatusNotFound, err.Error())
		return
	}

	if complaint.UserID != userID {
		utils.Error(w, http.StatusForbidden, "Forbidden")
		return
	}

	var req dto.UpdateComplaintRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	msg := validation.ValidateComplaint(req.Title, req.Description)
	if msg != "" {
		utils.Error(w, http.StatusBadRequest, msg)
		return
	}

	complaint.Title = req.Title
	complaint.Description = req.Description

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

// DeleteComplaint godoc
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

	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/complaints/"))
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid complaint ID")
		return
	}

	userID := r.Context().Value("userID").(int)

	complaint, err := complaintService.GetByID(id)
	if err != nil {
		utils.Error(w, http.StatusNotFound, err.Error())
		return
	}

	if complaint.UserID != userID {
		utils.Error(w, http.StatusForbidden, "Forbidden")
		return
	}

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
