package mapper

import (
	"gear-server/internal/domain"
)

type ProfileResponse struct {
	ID          string  `json:"id"`
	FullName    string  `json:"full_name"`
	PhoneNumber string  `json:"phone_number"`
	Address     *string `json:"address"`
	AvatarUrl   *string `json:"avatar_url"`
	UserId      string  `json:"user_id"`
}

func ToProfileResponse(profile *domain.Profile) *ProfileResponse {
	if profile == nil {
		return nil
	}
	return &ProfileResponse{
		ID:          profile.ID,
		FullName:    profile.FullName,
		PhoneNumber: profile.PhoneNumber,
		Address:     profile.Address,
		AvatarUrl:   profile.AvatarUrl,
		UserId:      profile.UserId,
	}
}

func ToProfileResponses(profiles []*domain.Profile) []*ProfileResponse {
	if profiles == nil {
		return nil
	}

	result := make([]*ProfileResponse, 0, len(profiles))

	for _, profile := range profiles {
		result = append(result, ToProfileResponse(profile))
	}

	return result
}
