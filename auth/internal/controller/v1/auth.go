package v1

import (
	"fmt"
	"net/http"

	"github.com/adslmgrv/mycourses-backend/auth/internal/domain"
	apperr "github.com/adslmgrv/mycourses-backend/auth/internal/error"
	"github.com/adslmgrv/mycourses-backend/auth/internal/service"
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

func (a AuthController) MakeRoutes(e *gin.Engine) {
	e.POST("/api/v1/auth/users", apperr.HandleAppErr(a.signUp))
}

func (a *AuthController) signUp(c *gin.Context) error {
	var request domain.SignUpRequest

	err := c.BindJSON(&request)

	if err != nil {
		fmt.Printf("%s", err)
		return err
	}

	err = a.authService.SignUp(c.Request.Context(), request)

	if err != nil {
		return err
	}

	c.Status(http.StatusCreated)

	return nil
}
