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

// Login godoc
//
// @Summary Login user
// @Description Login using email and password
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "Login Request"
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Failure 401 {object} utils.APIResponse
// @Router /login [post]
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

	response := dto.LoginResponse{
		User:  dto.ToUserResponse(user),
		Token: user.Token,
	}

	utils.Success(
		w,
		http.StatusOK,
		"Login successful",
		response,
	)
}

// ForgotPassword godoc
//
// @Summary Forgot Password
// @Description Send OTP to registered email
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body dto.ForgotPasswordRequest true "Forgot Password Request"
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Router /forgot-password [post]
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

// VerifyOTP godoc
//
// @Summary Verify OTP
// @Description Verify OTP sent to email
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body dto.VerifyOTPRequest true "Verify OTP Request"
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Router /verify-otp [post]
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

// ResetPassword godoc
//
// @Summary Reset Password
// @Description Reset password using OTP
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body dto.ResetPasswordRequest true "Reset Password Request"
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Router /reset-password [post]
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
