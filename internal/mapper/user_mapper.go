package mapper

import (
	"gear-server/internal/domain"
)

type UserResponse struct {
	ID       string           `json:"id"`
	Username string           `json:"username"`
	Email    string           `json:"email"`
	Profile  ProfileResponse `json:"profile"`
}

func ToUserResponse(user *domain.User) *UserResponse {
	if user == nil {
		return nil
	}
	return &UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Profile:  *ToProfileResponse(&user.Profile),
	}
}

func ToUserResponses(users []*domain.User) []*UserResponse {
	if users == nil {
		return nil
	}

	result := make([]*UserResponse, 0, len(users))

	for _, user := range users {
		result = append(result, ToUserResponse(user))
	}

	return result
}
