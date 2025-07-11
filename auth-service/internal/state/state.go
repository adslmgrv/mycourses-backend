package state

import "github.com/adslmgrv/mycourses-backend/auth-service/internal/repo"

type State struct {
	UserRepo repo.UserRepo
}
