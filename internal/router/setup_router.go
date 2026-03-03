package router

import (
	"gear-server/internal/handler"
	"gear-server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *handler.AuthHandler) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(middleware.ExceptionMiddleware())
	
	auth := r.Group("/auth")
	{
		auth.POST("/signin", userHandler.SignIn)
	}

	return r
}
