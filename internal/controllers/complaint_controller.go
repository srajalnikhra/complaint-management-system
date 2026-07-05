package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/srajalnikhra/complaint-management-system/internal/dto"
	"github.com/srajalnikhra/complaint-management-system/internal/models"
	"github.com/srajalnikhra/complaint-management-system/internal/services"
)

var complaintService = services.NewComplaintService()

func CreateComplaint(w http.ResponseWriter, r *http.Request) {

	var req dto.CreateComplaintRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	userIDStr := r.Context().Value("userID").(int)

	userID := userIDStr
	fmt.Println("User ID From Context:", userIDStr)
	fmt.Println("Converted User ID:", userID)

	complaint := models.Complaint{
		UserID:      userID,
		Title:       req.Title,
		Description: req.Description,
	}

	if err := complaintService.Create(&complaint); err != nil {
		fmt.Println("Create Complaint Error:", err)

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(complaint)

}

func GetMyComplaints(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value("userID").(int)

	complaints, err := complaintService.GetByUserID(userID)
	if err != nil {
		http.Error(w, "Failed to fetch complaints", http.StatusInternalServerError)
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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetComplaintByID(w http.ResponseWriter, r *http.Request) {

	idStr := strings.TrimPrefix(r.URL.Path, "/complaints/")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid complaint id", http.StatusBadRequest)
		return
	}

	complaint, err := complaintService.GetByID(id)
	if err != nil {
		http.Error(w, "Complaint not found", http.StatusNotFound)
		return
	}

	userID := r.Context().Value("userID").(int)

	if complaint.UserID != userID {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	response := dto.ComplaintResponse{
		ID:          complaint.ID,
		Title:       complaint.Title,
		Description: complaint.Description,
		Status:      complaint.Status,
		CreatedAt:   complaint.CreatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func UpdateComplaint(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/complaints/"))
	if err != nil {
		http.Error(w, "Invalid complaint ID", http.StatusBadRequest)
		return
	}

	userID := r.Context().Value("userID").(int)

	complaint, err := complaintService.GetByID(id)
	if err != nil {
		http.Error(w, "Complaint not found", http.StatusNotFound)
		return
	}

	if complaint.UserID != userID {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	var req dto.UpdateComplaintRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	complaint.Title = req.Title
	complaint.Description = req.Description

	if err := complaintService.Update(complaint); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(complaint)
}

func DeleteComplaint(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/complaints/"))
	if err != nil {
		http.Error(w, "Invalid complaint ID", http.StatusBadRequest)
		return
	}

	userID := r.Context().Value("userID").(int)

	complaint, err := complaintService.GetByID(id)
	if err != nil {
		http.Error(w, "Complaint not found", http.StatusNotFound)
		return
	}

	if complaint.UserID != userID {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	if err := complaintService.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Complaint deleted successfully"))
}