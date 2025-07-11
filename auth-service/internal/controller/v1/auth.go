package controller

import (
	"github.com/adslmgrv/mycourses-backend/auth-service/internal/service"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return AuthController{
		authService: authService,
	}
}

// /api/v1/auth/signup
func (_c *AuthController) SignUp(c *gin.Context) error {
	// convert body to json dto dto.SignUpRequest
	// var request dto.SignUpRequest

	// err := c.BindJSON(&request)

	// if err != nil {
	// 	return appe.Errorf(appe.BadRequestError, "Invalid request body")
	// }
	return nil
}
