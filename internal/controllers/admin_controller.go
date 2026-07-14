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

// GetAllComplaints godoc
//
// @Summary Get All Complaints
// @Description Get all complaints with pagination, filtering and sorting (Admin only)
// @Tags Admin
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page Number"
// @Param limit query int false "Items Per Page"
// @Param search query string false "Search by title"
// @Param status query string false "Complaint Status"
// @Param sort query string false "Sort Field"
// @Param order query string false "ASC or DESC"
// @Success 200 {object} utils.APIResponse
// @Failure 401 {object} utils.APIResponse
// @Failure 403 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /admin/complaints [get]
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

// UpdateComplaintStatus godoc
//
// @Summary Update Complaint Status
// @Description Update complaint status (Admin only)
// @Tags Admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Complaint ID"
// @Param request body dto.UpdateComplaintStatusRequest true "Status"
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Failure 401 {object} utils.APIResponse
// @Failure 403 {object} utils.APIResponse
// @Failure 404 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /admin/complaints/{id}/status [patch]
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

		if err == utils.ErrComplaintNotFound {
			utils.Error(w, http.StatusNotFound, err.Error())
			return
		}

		utils.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(w, http.StatusOK, "Complaint status updated successfully", nil)
}

// GetAllUsers godoc
//
// @Summary Get All Users
// @Description Get all users (Admin only)
// @Tags Admin
// @Produce json
// @Security BearerAuth
// @Success 200 {object} utils.APIResponse
// @Failure 401 {object} utils.APIResponse
// @Failure 403 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /admin/users [get]
func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	users, err := adminService.GetAllUsers()
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(
		w,
		http.StatusOK,
		"Users fetched successfully",
		users,
	)
}

// UpdateUserRole godoc
//
// @Summary Update User Role
// @Description Update user role (Admin only)
// @Tags Admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "User ID"
// @Param request body dto.UpdateUserRoleRequest true "Role"
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Failure 401 {object} utils.APIResponse
// @Failure 403 {object} utils.APIResponse
// @Failure 404 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /admin/users/{id}/role [patch]
func UpdateUserRole(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(
		strings.TrimSuffix(
			strings.TrimPrefix(r.URL.Path, "/admin/users/"),
			"/role",
		),
	)

	if err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var req dto.UpdateUserRoleRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if msg := validation.ValidateUserRole(req.Role); msg != "" {
		utils.Error(w, http.StatusBadRequest, msg)
		return
	}

	if err := adminService.UpdateUserRole(id, req.Role); err != nil {
		if err == utils.ErrUserNotFound {
			utils.Error(w, http.StatusNotFound, err.Error())
			return
		}

		utils.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(
		w,
		http.StatusOK,
		"User role updated successfully",
		nil,
	)
}

// UpdateUserStatus godoc
//
// @Summary Update User Status
// @Description Activate or deactivate a user (Admin only)
// @Tags Admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "User ID"
// @Param request body dto.UpdateUserStatusRequest true "User Status"
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Failure 401 {object} utils.APIResponse
// @Failure 403 {object} utils.APIResponse
// @Failure 404 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /admin/users/{id}/status [patch]
func UpdateUserStatus(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(
		strings.TrimSuffix(
			strings.TrimPrefix(r.URL.Path, "/admin/users/"),
			"/status",
		),
	)

	if err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var req dto.UpdateUserStatusRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := adminService.UpdateUserStatus(id, req.IsActive); err != nil {
		if err == utils.ErrUserNotFound {
			utils.Error(w, http.StatusNotFound, err.Error())
			return
		}

		utils.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(
		w,
		http.StatusOK,
		"User status updated successfully",
		nil,
	)
}

// DeleteUser godoc
//
// @Summary Delete User
// @Description Delete a user by ID (Admin only)
// @Tags Admin
// @Produce json
// @Security BearerAuth
// @Param id path int true "User ID"
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Failure 401 {object} utils.APIResponse
// @Failure 403 {object} utils.APIResponse
// @Failure 404 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /admin/users/{id} [delete]
func DeleteUser(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, "/admin/users/"), ""))
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	adminID := r.Context().Value("userID").(int)

	if id == adminID {
		utils.Error(
			w,
			http.StatusBadRequest,
			"You cannot delete your own account",
		)
		return
	}

	if err := adminService.DeleteUser(id); err != nil {

		if err == utils.ErrUserNotFound {
			utils.Error(
				w,
				http.StatusNotFound,
				err.Error(),
			)
			return
		}

		utils.Error(
			w,
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	utils.Success(
		w,
		http.StatusOK,
		"User deleted successfully",
		nil,
	)
}
