package v1

import (
	"net/http"

	"github.com/adslmgrv/mycourses-backend/auth-service/internal/domain"
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

func (a *AuthController) SignUp(c *gin.Context) {
	var request domain.SignUpRequest

	c.BindJSON(&request)

	// if err != nil {
	// 	return appe.Errorf(appe.BadRequestError, "Invalid request body")
	// }

	a.authService.SignUp(c.Request.Context(), request)

	// if err != nil {
	// 	return err
	// }

	c.Status(http.StatusCreated)

	// return nil
}
