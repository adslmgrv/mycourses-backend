package repository

import (
	"context"

	"github.com/adslmgrv/mycourses-backend/auth/internal/model"
)

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	UpdatePasswordHashByEmail(ctx context.Context, email string, passwordHash []byte) error
	CreateUser(ctx context.Context, user model.User) error
}
