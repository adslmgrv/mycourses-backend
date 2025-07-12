package error

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleAppErr(f func(c *gin.Context) error) func(c *gin.Context) {
	return func(c *gin.Context) {
		err := f(c)

		if err != nil {
			status, kind, message := func() (int, int, string) {
				if appErr, ok := err.(AppError); ok {
					return appErr.Kind().httpStatus(), int(appErr.Kind()), appErr.Message()
				} else if e, ok := err.(gin.Error); ok {
					return http.StatusBadRequest, 1, e.Error()
				} else {
					log.Printf("Cause of internal server error: %s", err)
					return http.StatusInternalServerError, 0, "Internal server error"
				}
			}()

			c.AbortWithStatusJSON(status, map[string]any{
				"message": message,
				"kind":    kind,
				"success": false,
				"data":    nil,
			})
		}
	}
}

func (k AppErrorKind) httpStatus() int {
	switch k {
	case InvalidCredentialsError:
		return http.StatusUnauthorized
	default:
		return http.StatusBadRequest
	}
}
