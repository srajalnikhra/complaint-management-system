package services

import (
	"errors"
	"time"

	"github.com/srajalnikhra/complaint-management-system/internal/dto"
	"github.com/srajalnikhra/complaint-management-system/internal/models"
	"github.com/srajalnikhra/complaint-management-system/internal/repositories"
	"github.com/srajalnikhra/complaint-management-system/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

// LoginService coordinates authentication operations handling validation and tokens.
type LoginService struct {
	repo *repositories.LoginRepository
}

// NewLoginService creates a new instance of the LoginService.
func NewLoginService() *LoginService {
	return &LoginService{
		repo: repositories.NewLoginRepository(),
	}
}

// Login validates user credentials and returns user details with a JWT token.
func (s *LoginService) Login(req dto.LoginRequest) (*models.User, error) {

	// Get active user data from the database.
	user, err := s.repo.GetByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Compare the entered password with stored bcrypt password hash.
	if !utils.CheckPassword(req.Password, user.Password) {
		return nil, errors.New("invalid email or password")
	}

	// Check if the user is active.
	if !user.IsActive {
		return nil, errors.New("your account has been deactivated")
	}

	// Generate a secure JWT access token.
	token, err := utils.GenerateJWT(user)
	if err != nil {
		return nil, err
	}

	user.Password = ""

	user.Token = token

	return user, nil
}

// SendForgotPasswordOTP generates and emails a temporary OTP code for password resets.
func (s *LoginService) SendForgotPasswordOTP(email string) error {

	// Check if user exists matching target email.
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return errors.New("user not found")
	}

	// Create and hash a random verification OTP.
	otp := utils.GenerateOTP()

	hash, err := bcrypt.GenerateFromPassword([]byte(otp), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	expiry := time.Now().Add(10 * time.Minute)

	// Write the OTP hash and its expiry timestamp to the user record.
	err = s.repo.SaveOTP(
		email,
		string(hash),
		expiry,
	)
	if err != nil {
		return err
	}

	// Deliver the OTP to user via email.
	err = SendOTPEmail(
		user.Email,
		user.Name,
		otp,
	)

	if err != nil {
		return err
	}

	return nil
}

// VerifyOTP validates the OTP code input from user.
func (s *LoginService) VerifyOTP(req dto.VerifyOTPRequest) error {

	user, err := s.repo.GetOTPData(req.Email)
	if err != nil {
		return errors.New("invalid email")
	}

	if user.OTPHash == "" {
		return errors.New("otp not found")
	}

	if user.OTPExpiresAt == nil {
		return errors.New("otp not found")
	}

	// Check validation timestamps against the expiration time.
	if time.Now().After(*user.OTPExpiresAt) {
		return errors.New("otp has expired")
	}

	// Compare entered OTP with the stored OTP hash.
	err = bcrypt.CompareHashAndPassword(
		[]byte(user.OTPHash),
		[]byte(req.OTP),
	)

	if err != nil {
		return errors.New("invalid otp")
	}

	return nil
}

// ResetPassword overrides the old password with the newly submitted choice.
func (s *LoginService) ResetPassword(req dto.ResetPasswordRequest) error {

	// Check user OTP validity first.
	user, err := s.repo.GetOTPData(req.Email)
	if err != nil {
		return errors.New("invalid email")
	}

	if user.OTPHash == "" {
		return errors.New("otp not found")
	}

	if user.OTPExpiresAt == nil {
		return errors.New("otp not found")
	}

	if time.Now().After(*user.OTPExpiresAt) {
		return errors.New("otp has expired")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.OTPHash),
		[]byte(req.OTP),
	)

	if err != nil {
		return errors.New("invalid otp")
	}

	// Hash the new password value.
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.NewPassword),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	// Update the password value in the database.
	err = s.repo.UpdatePassword(
		req.Email,
		string(hashedPassword),
	)

	if err != nil {
		return err
	}

	// Clear the OTP fields to prevent reuse.
	err = s.repo.ClearOTP(req.Email)
	if err != nil {
		return err
	}

	return nil
}
