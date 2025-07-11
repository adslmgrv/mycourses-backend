package repo

import (
	"context"

	"github.com/adslmgrv/mycourses-backend/auth-service/error"
	"github.com/adslmgrv/mycourses-backend/auth-service/model"
)

type UserRepo interface {
	FindByUsername(ctx context.Context, username string) (*model.User, error.Error)
	CreateUser(ctx context.Context, user model.User) error.Error
}
