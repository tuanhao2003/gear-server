package serviceimpl

import (
	"errors"
	"gear-server/internal/domain"
	"gear-server/internal/repository"

	"gear-server/internal/service"
	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) service.AuthService {
	return &authService{
		userRepo: userRepo,
	}
}

func (s *authService) SignIn(userNameOrEmail string, password string) (*domain.User, error) {
	var user *domain.User
	var err error

	user, err = s.userRepo.FindByUsername(userNameOrEmail)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			user, err = s.userRepo.FindByEmail(userNameOrEmail)
			if err != nil {
				return nil, errors.New("invalid credentials")
			}
		} else {
			return nil, err
		}
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.HashPassword),
		[]byte(password),
	); err != nil {
		return nil, errors.New("invalid password")
	}

	return user, nil
}
