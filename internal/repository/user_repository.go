package repository

import "gear-server/internal/domain"

type UserRepository interface {
	FindByUsername(username string) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
}
