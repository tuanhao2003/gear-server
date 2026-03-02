package router

import (
	"gear-server/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *handler.AuthHandler) *gin.Engine {
	r := gin.Default()

	auth := r.Group("/auth")
	{
		auth.POST("/signin", userHandler.SignIn)
	}

	return r
}
