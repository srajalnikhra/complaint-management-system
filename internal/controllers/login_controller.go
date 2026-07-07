package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/srajalnikhra/complaint-management-system/internal/dto"
	"github.com/srajalnikhra/complaint-management-system/internal/services"
	"github.com/srajalnikhra/complaint-management-system/internal/utils"
	"github.com/srajalnikhra/complaint-management-system/internal/validation"
)

type LoginController struct {
	service *services.LoginService
}

func NewLoginController() *LoginController {
	return &LoginController{
		service: services.NewLoginService(),
	}
}

func (c *LoginController) Login(w http.ResponseWriter, r *http.Request) {

	var req dto.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	msg := validation.ValidateLogin(
		req.Email,
		req.Password,
	)

	if msg != "" {
		utils.Error(w, http.StatusBadRequest, msg)
		return
	}

	user, err := c.service.Login(req)
	if err != nil {
		utils.Error(w, http.StatusUnauthorized, err.Error())
		return
	}

	utils.Success(
		w,
		http.StatusOK,
		"Login successful",
		user,
	)
}