package service

import "gear-server/internal/domain"

type AuthService interface {
	SignIn(userNameOrEmail string, password string) (*domain.User, error)
}