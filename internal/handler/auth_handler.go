package handler

import (
	"net/http"

	"gear-server/internal/dto"
	"gear-server/internal/enum"
	"gear-server/internal/mapper"
	"gear-server/internal/service"
	"gear-server/pkg/helper"

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
func (h *AuthHandler) SignIn(c *gin.Context) {
	var requestDto dto.SignInRequest

	if err := c.ShouldBind(&requestDto); err != nil {
		c.JSON(http.StatusBadRequest,
			helper.Response(
				enum.VALIDATION_ERROR,
				"auth validation error",
				nil,
				helper.WithError[error](err),
			),
		)
		return
	}

	usernameOrEmail := requestDto.Username
	if usernameOrEmail == "" {
		usernameOrEmail = requestDto.Email
	}

	userEntity, err := h.authService.SignIn(usernameOrEmail, requestDto.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized,
			helper.Response(
				enum.UNAUTHORIZED,
				"Invalid credentials",
				nil,
				helper.WithError[error](err),
			),
		)
		return
	}

	userResponse := *mapper.ToUserResponse(userEntity)

	c.JSON(http.StatusOK,
		helper.Response(
			enum.SUCCESS,
			"success",
			userResponse,
		),
	)
}
