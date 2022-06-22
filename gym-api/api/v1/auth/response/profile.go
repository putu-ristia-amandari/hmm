package response

import (
	"gym-membership/business/user"
	"gym-membership/helpers"
)

type ProfileResponse struct {
	ID               int                 `json:"id"`
	Membership       *MembershipResponse `json:"membership"`
	Name             string              `json:"name"`
	Gender           string              `json:"gender"`
	BirthOfDate      string              `json:"birth_of_date"`
	Height           int                 `json:"height"`
	Weight           int                 `json:"Weight"`
	Email            string              `json:"email"`
	Handphone        string              `json:"handphone"`
	Address          string              `json:"address"`
	City             string              `json:"city"`
	Province         string              `json:"province"`
	Nationality      string              `json:"nationality"`
	Photo            string              `json:"photo"`
	Active           bool                `json:"active"`
	ActiveMembership bool                `json:"active_membership"`
	VerifiedAt       string              `json:"verified_at"`
	CreatedAt        string              `json:"created_at"`
	UpdatedAt        string              `json:"updated_at"`
}

type MembershipResponse struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Image          string `json:"image"`
	LimitedClass   int    `json:"limited_class"`
	LimitedTime    int    `json:"limited_time"`
	LimitedContent int    `json:"limited_content"`
	Description    string `json:"description"`
	Details        string `json:"details"`
}

func CreateProfileResponse(user *user.User) *ProfileResponse {
	return &ProfileResponse{
		ID: user.ID,
		Membership: &MembershipResponse{
			ID:             user.Membership.ID,
			Name:           user.Membership.Name,
			Image:          user.Membership.Image,
			LimitedClass:   user.Membership.LimitedClass,
			LimitedTime:    user.Membership.LimitedTime,
			LimitedContent: user.Membership.LimitedContent,
			Description:    user.Membership.Description,
			Details:        user.Membership.Details,
		},
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
		Active:           user.Status,
		ActiveMembership: user.StatusMembership,
		VerifiedAt:       helpers.TimeFormat(user.VerifiedAt),
		CreatedAt:        helpers.TimeFormat(user.CreatedAt),
		UpdatedAt:        helpers.TimeFormat(user.UpdatedAt),
	}
}
