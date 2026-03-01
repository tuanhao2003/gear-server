package service

import (
	"errors"
	"gear-server/internal/domain"
	"gear-server/internal/repository"

	// "golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) SignIn(userNameOrEmail string, password string) (*domain.User, error) {
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

	// if err := bcrypt.CompareHashAndPassword(
	// 	[]byte(user.HashPassword),
	// 	[]byte(password),
	// ); err != nil {
	// 	return nil, errors.New("invalid credentials")
	// }

	if user.HashPassword != password {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}
