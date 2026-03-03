package middleware

import (
	"fmt"
	"net/http"

	"gear-server/internal/enum"
	"gear-server/pkg/helper"

	"github.com/gin-gonic/gin"
)

func ExceptionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		defer func() {
			if r := recover(); r != nil {
				err := fmt.Errorf("%v", r)

				c.AbortWithStatusJSON(
					http.StatusInternalServerError,
					helper.Response(
						enum.INTERNAL_SERVER_ERROR,
						"internal server error",
						nil,
						helper.WithError[error](err),
					),
				)
			}
		}()

		c.Next()
	}
}
