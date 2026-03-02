package handler

import (
	"net/http"

	"gear-server/internal/dto"
	"gear-server/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) SignIn(context *gin.Context) {
	var requestDto dto.SignInRequest

	if err := context.ShouldBind(&requestDto); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	usernameOrEmail := requestDto.Username
	if usernameOrEmail == "" {
		usernameOrEmail = requestDto.Email
	}

	user, err := h.authService.SignIn(usernameOrEmail, requestDto.Password)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, user)
}
