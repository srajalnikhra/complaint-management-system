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
		Role: models.RoleUser,
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
