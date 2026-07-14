package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/srajalnikhra/complaint-management-system/internal/dto"
	"github.com/srajalnikhra/complaint-management-system/internal/models"
	"github.com/srajalnikhra/complaint-management-system/internal/services"
	"github.com/srajalnikhra/complaint-management-system/internal/utils"
	"github.com/srajalnikhra/complaint-management-system/internal/validation"
)

type UserController struct {
	service *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		service: services.NewUserService(),
	}
}

// Register godoc
//
// @Summary Register a new user
// @Description Create a new user account
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body dto.RegisterUserRequest true "Register Request"
// @Success 201 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Failure 409 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /register [post]
func (c *UserController) Register(w http.ResponseWriter, r *http.Request) {

	var req dto.RegisterUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	msg := validation.ValidateRegister(
		req.Name,
		req.Email,
		req.Password,
	)

	if msg != "" {
		utils.Error(w, http.StatusBadRequest, msg)
		return
	}

	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Role:     models.RoleUser,
	}

	if err := c.service.CreateUser(&user); err != nil {

		if err == utils.ErrUserAlreadyExists {
			utils.Error(w, http.StatusConflict, err.Error())
			return
		}

		utils.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(
		w,
		http.StatusCreated,
		"User registered successfully",
		user,
	)
}
