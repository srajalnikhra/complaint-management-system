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

func (c *LoginController) ForgotPassword(w http.ResponseWriter, r *http.Request) {

	var req dto.ForgotPasswordRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	msg := validation.ValidateForgotPassword(req.Email)

	if msg != "" {
		utils.Error(w, http.StatusBadRequest, msg)
		return
	}

	err := c.service.SendForgotPasswordOTP(req.Email)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.Success(
		w,
		http.StatusOK,
		"OTP sent successfully",
		nil,
	)
}

func (c *LoginController) VerifyOTP(w http.ResponseWriter, r *http.Request) {

	var req dto.VerifyOTPRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if msg := validation.ValidateVerifyOTP(req.Email, req.OTP); msg != "" {
		utils.Error(w, http.StatusBadRequest, msg)
		return
	}

	err := c.service.VerifyOTP(req)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.Success(
		w,
		http.StatusOK,
		"OTP verified successfully",
		nil,
	)
}

func (c *LoginController) ResetPassword(w http.ResponseWriter, r *http.Request) {

	var req dto.ResetPasswordRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if msg := validation.ValidateResetPassword(
		req.Email,
		req.OTP,
		req.NewPassword,
	); msg != "" {
		utils.Error(w, http.StatusBadRequest, msg)
		return
	}

	err := c.service.ResetPassword(req)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.Success(
		w,
		http.StatusOK,
		"Password reset successfully",
		nil,
	)
}
