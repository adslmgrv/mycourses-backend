package service

import "github.com/adslmgrv/mycourses-backend/auth-service/internal/repository"

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return UserService{
		repo: repo,
	}
}
