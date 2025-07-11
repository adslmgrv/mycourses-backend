package state

import "github.com/adslmgrv/mycourses-backend/auth-service/repo"

type State struct {
	UserRepo repo.UserRepo
}
