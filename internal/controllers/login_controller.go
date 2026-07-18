package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/srajalnikhra/complaint-management-system/internal/dto"
	"github.com/srajalnikhra/complaint-management-system/internal/services"
	"github.com/srajalnikhra/complaint-management-system/internal/utils"
	"github.com/srajalnikhra/complaint-management-system/internal/validation"
)

// LoginController handles authentication-related HTTP endpoints.
type LoginController struct {
	service *services.LoginService
}

// NewLoginController creates a new instance of the LoginController.
func NewLoginController() *LoginController {
	return &LoginController{
		service: services.NewLoginService(),
	}
}

// Login authenticates a user and returns their profile with an access token.
// If the credentials are valid, it creates a JWT token and returns it.
//
// # Login godoc
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

	// Decode the request body to get the login credentials.
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Validate the login request parameters.
	msg := validation.ValidateLogin(
		req.Email,
		req.Password,
	)

	if msg != "" {
		utils.Error(w, http.StatusBadRequest, msg)
		return
	}

	// Authenticate the user and generate a JWT token.
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

// ForgotPassword initiates the password recovery flow.
// It generates an OTP and emails it to the requested email address if the user exists.
//
// # ForgotPassword godoc
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

	// Decode the request body to get the user's email.
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Validate the email address format.
	msg := validation.ValidateForgotPassword(req.Email)

	if msg != "" {
		utils.Error(w, http.StatusBadRequest, msg)
		return
	}

	// Create a password reset OTP and email it to the user.
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

// VerifyOTP checks if the submitted OTP is correct and has not expired.
//
// # VerifyOTP godoc
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

	// Decode the request body to get the email and OTP.
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Validate OTP and email format.
	if msg := validation.ValidateVerifyOTP(req.Email, req.OTP); msg != "" {
		utils.Error(w, http.StatusBadRequest, msg)
		return
	}

	// Validate the OTP against the database.
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

// ResetPassword updates the user's password to a new one after verifying the OTP.
//
// # ResetPassword godoc
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

	// Decode the request body to get the new password and OTP.
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Validate input parameters format.
	if msg := validation.ValidateResetPassword(
		req.Email,
		req.OTP,
		req.NewPassword,
	); msg != "" {
		utils.Error(w, http.StatusBadRequest, msg)
		return
	}

	// Verify the OTP one last time and save the new hashed password.
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
