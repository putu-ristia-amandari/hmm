package response

import (
	"gym-membership/business/user"
	"gym-membership/helpers"
)

type RegisterResponse struct {
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	Email     string `json:"email"`
	Handphone string `json:"handphone"`
	Active    bool   `json:"active"`
	CreatedAt string `json:"created_at"`
}

func CreateRegisterResponse(user *user.User) *RegisterResponse {
	return &RegisterResponse{
		Name:      user.Name,
		Gender:    user.Gender,
		Email:     user.Email,
		Handphone: user.Handphone,
		Active:    user.Status,
		CreatedAt: helpers.TimeFormat(user.CreatedAt),
	}
}
