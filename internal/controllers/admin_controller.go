package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/srajalnikhra/complaint-management-system/internal/dto"
	"github.com/srajalnikhra/complaint-management-system/internal/services"
)

var adminService = services.NewAdminService()

func GetAllComplaints(w http.ResponseWriter, r *http.Request) {

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	search := r.URL.Query().Get("search")
	status := r.URL.Query().Get("status")

	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	complaints, err := adminService.GetAllComplaints(page, limit, search, status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(complaints)
}

func UpdateComplaintStatus(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, "/admin/complaints/"), "/status"))
	if err != nil {
		http.Error(w, "Invalid complaint ID", http.StatusBadRequest)
		return
	}

	var req dto.UpdateComplaintStatusRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	if err := adminService.UpdateComplaintStatus(id, req.Status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Complaint status updated successfully"))
}
