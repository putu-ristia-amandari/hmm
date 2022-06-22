package response

import (
	"gym-membership/business/user"
	"gym-membership/helpers"
)

type LoginResponse struct {
	User  *UserResponse `json:"user"`
	Token string        `json:"token"`
}

type UserResponse struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Gender           string `json:"gender"`
	BirthOfDate      string `json:"birth_of_date"`
	Height           int    `json:"height"`
	Weight           int    `json:"Weight"`
	Email            string `json:"email"`
	Handphone        string `json:"handphone"`
	Address          string `json:"address"`
	City             string `json:"city"`
	Province         string `json:"province"`
	Nationality      string `json:"nationality"`
	Photo            string `json:"photo"`
	Status           bool   `json:"status"`
	StatusMembership bool   `json:"status_membership"`
	VerifiedAt       string `json:"verified_at"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}

func CreateLoginResponse(user *user.User, token string) *LoginResponse {
	return &LoginResponse{
		User: &UserResponse{
			Name:             user.Name,
			Email:            user.Email,
			Handphone:        user.Handphone,
			Address:          user.Address,
			City:             user.City,
			Province:         user.Province,
			Nationality:      user.Nationality,
			Gender:           user.Gender,
			BirthOfDate:      helpers.DateFormat(user.BirthOfDate),
			Height:           user.Height,
			Weight:           user.Weight,
			Photo:            user.Photo,
			Status:           user.Status,
			StatusMembership: user.StatusMembership,
			VerifiedAt:       helpers.TimeFormat(user.VerifiedAt),
			CreatedAt:        helpers.TimeFormat(user.CreatedAt),
			UpdatedAt:        helpers.TimeFormat(user.UpdatedAt),
		},
		Token: token,
	}
}
