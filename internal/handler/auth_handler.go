package handler

import (
	"net/http"

	"gear-server/internal/dto"
	"gear-server/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	authService service.AuthService
}

func NewUserHandler(authService service.AuthService) *UserHandler {
	return &UserHandler{
		authService: authService,
	}
}

func (h *UserHandler) SignIn(context *gin.Context) {
	var requestDto dto.SignInRequest

	if err := context.ShouldBindJSON(&requestDto); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	user, err := h.authService.SignIn(requestDto.UsernameOrEmail, requestDto.Password)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, user)
}
