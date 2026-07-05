package services

import (
	"errors"

	"github.com/srajalnikhra/complaint-management-system/internal/dto"
	"github.com/srajalnikhra/complaint-management-system/internal/models"
	"github.com/srajalnikhra/complaint-management-system/internal/repositories"
	"github.com/srajalnikhra/complaint-management-system/internal/utils"
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

	token, err := utils.GenerateJWT(user)
	if err != nil {
		return nil, err
	}

	user.Password = ""

	user.Token = token

	return user, nil
}
