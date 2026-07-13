package service

import (
    "errors"

    "github.com/google/uuid"

    "github.com/adullahkapadia/auth-service/internal/model"
    "github.com/adullahkapadia/auth-service/internal/repository"
    "github.com/adullahkapadia/auth-service/internal/utils"
)

type AuthService struct {
    Repo *repository.UserRepository
}


func (s *AuthService) Register(req *model.RegisterRequest) error {

    exists, err := s.Repo.Exists(req.Email)

    if err != nil {
        return err
    }

    if exists {
        return errors.New("email already exists")
    }

    hash, err := utils.HashPassword(req.Password)

    if err != nil {
        return err
    }

    user := &model.User{
        ID:       uuid.New().String(),
        Name:     req.Name,
        Email:    req.Email,
        Password: hash,
    }

    return s.Repo.Create(user)
}


func (s *AuthService) Login(req *model.LoginRequest, secret string) (string, error) {

	user, err := s.Repo.GetByEmail(req.Email)

	if err != nil {
		return "", err
	}

	err = utils.CheckPassword(
		user.Password,
		req.Password,
	)

	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(
		user.ID,
		secret,
	)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) GetProfile(userID string) (*model.User, error) {
	return s.Repo.GetByID(userID)
}