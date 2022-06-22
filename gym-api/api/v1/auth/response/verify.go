package response

import (
	"gym-membership/helpers"
	"time"
)

type VerifyResponse struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	VerifiedAt string `json:"verified_at"`
}

func CreateVerifyResponse(name string, email string, verifiedAt time.Time) *VerifyResponse {
	return &VerifyResponse{
		Name:       name,
		Email:      email,
		VerifiedAt: helpers.TimeFormat(verifiedAt),
	}
}
