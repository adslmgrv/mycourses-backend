package controller

import (
	"net/http"

	"github.com/adslmgrv/mycourses-backend/auth-service/internal/domain"
	appe "github.com/adslmgrv/mycourses-backend/auth-service/internal/error"
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

func (a *AuthController) SignUp(c *gin.Context) error {
	var request domain.SignUpRequest

	err := c.BindJSON(&request)

	if err != nil {
		return appe.Errorf(appe.BadRequestError, "Invalid request body")
	}

	err = a.authService.SignUp(c.Request.Context(), request)

	if err != nil {
		return err
	}

	c.Status(http.StatusCreated)

	return nil
}
