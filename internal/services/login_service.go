package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/srajalnikhra/complaint-management-system/internal/dto"
	"github.com/srajalnikhra/complaint-management-system/internal/models"
	"github.com/srajalnikhra/complaint-management-system/internal/repositories"
	"github.com/srajalnikhra/complaint-management-system/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	repo *repositories.LoginRepository
}

func NewLoginService() *LoginService {
	return &LoginService{
		repo: repositories.NewLoginRepository(),
	}
}

func (s *LoginService) Login(req dto.LoginRequest) (*models.User, error) {

	user, err := s.repo.GetByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if !utils.CheckPassword(req.Password, user.Password) {
		return nil, errors.New("invalid email or password")
	}

	if !user.IsActive {
		return nil, errors.New("your account has been deactivated")
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		return nil, err
	}

	user.Password = ""

	user.Token = token

	return user, nil
}

func (s *LoginService) SendForgotPasswordOTP(email string) error {

	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return errors.New("user not found")
	}

	otp := utils.GenerateOTP()

	hash, err := bcrypt.GenerateFromPassword([]byte(otp), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	expiry := time.Now().Add(10 * time.Minute)

	err = s.repo.SaveOTP(
		email,
		string(hash),
		expiry,
	)
	if err != nil {
		return err
	}

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

func (s *LoginService) VerifyOTP(req dto.VerifyOTPRequest) error {

	user, err := s.repo.GetOTPData(req.Email)
	if err != nil {
		return errors.New("invalid email")
	}

	if user.OTPHash == "" {
		return errors.New("otp not found")
	}

	fmt.Println("Current Time :", time.Now())
	fmt.Println("Expiry Time  :", *user.OTPExpiresAt)
	fmt.Println("Expired      :", time.Now().After(*user.OTPExpiresAt))

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

	return nil
}

func (s *LoginService) ResetPassword(req dto.ResetPasswordRequest) error {

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

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.NewPassword),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	err = s.repo.UpdatePassword(
		req.Email,
		string(hashedPassword),
	)

	if err != nil {
		return err
	}

	err = s.repo.ClearOTP(req.Email)
	if err != nil {
		return err
	}

	return nil
}

