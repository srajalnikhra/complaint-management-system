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
