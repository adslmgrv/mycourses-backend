package service

import (
	"context"
	"time"

	"github.com/adslmgrv/mycourses-backend/auth-service/dto"
	"github.com/adslmgrv/mycourses-backend/auth-service/error"
	"github.com/adslmgrv/mycourses-backend/auth-service/model"
	"github.com/adslmgrv/mycourses-backend/auth-service/repo"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo repo.UserRepo
}

func NewAuthService(repo repo.UserRepo) AuthService {
	return AuthService{
		repo: repo,
	}
}

func (s AuthService) SignUp(ctx context.Context, request dto.SignUpRequest) error.Error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return error.Errorf(error.InternalError, "failed to hash password, cause: %s", err)
	}

	return s.repo.CreateUser(ctx, model.User{
		Name:         request.Name,
		Email:        request.Name,
		PasswordHash: passwordHash,
		CreatedAt:    time.Now().UTC(),
	})
}
