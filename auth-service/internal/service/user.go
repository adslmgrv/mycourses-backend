package service

import "github.com/adslmgrv/mycourses-backend/auth-service/internal/repo"

type UserService struct {
	repo repo.UserRepo
}

func NewUserService(repo repo.UserRepo) UserService {
	return UserService{
		repo: repo,
	}
}
