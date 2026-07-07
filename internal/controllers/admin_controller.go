package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/srajalnikhra/complaint-management-system/internal/dto"
	"github.com/srajalnikhra/complaint-management-system/internal/services"
	"github.com/srajalnikhra/complaint-management-system/internal/utils"
	"github.com/srajalnikhra/complaint-management-system/internal/validation"
)

var adminService = services.NewAdminService()

func GetAllComplaints(w http.ResponseWriter, r *http.Request) {

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	search := r.URL.Query().Get("search")
	status := r.URL.Query().Get("status")

	sort := r.URL.Query().Get("sort")
	order := strings.ToUpper(r.URL.Query().Get("order"))

	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	if order == "" {
		order = "DESC"
	}

	complaints, err := adminService.GetAllComplaints(page, limit, search, status, sort, order)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(
		w,
		http.StatusOK,
		"Complaints fetched successfully",
		complaints,
	)
}

func UpdateComplaintStatus(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, "/admin/complaints/"), "/status"))
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid complaint ID")
		return
	}

	var req dto.UpdateComplaintStatusRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	msg := validation.ValidateComplaintStatus(req.Status)

	if msg != "" {
		utils.Error(w, http.StatusBadRequest, msg)
		return
	}

	if err := adminService.UpdateComplaintStatus(id, req.Status); err != nil {
		utils.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(
		w,
		http.StatusOK,
		"Complaint status updated successfully",
		nil,
	)
}