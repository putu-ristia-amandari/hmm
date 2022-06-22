package response

import (
	"gym-membership/business/membership"
	"gym-membership/helpers"
)

type GetMembershipByIDResponse struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	Image          string  `json:"image"`
	LimitedClass   int     `json:"limited_class"`
	LimitedTime    int     `json:"limited_time"`
	LimitedContent int     `json:"limited_content"`
	Description    string  `json:"description"`
	Details        string  `json:"details"`
	Price          float64 `json:"price"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
}

func CreateGetAllMemberResponse(memberships []*membership.Membership) []*GetMembershipByIDResponse {
	var membershipsResponse []*GetMembershipByIDResponse

	for _, membership := range memberships {
		membershipsResponse = append(membershipsResponse, CreateGetMembershipByIDResponse(membership))
	}

	return membershipsResponse
}

func CreateGetMembershipByIDResponse(membership *membership.Membership) *GetMembershipByIDResponse {
	return &GetMembershipByIDResponse{
		ID:             membership.ID,
		Name:           membership.Name,
		Image:          membership.Image,
		LimitedClass:   membership.LimitedClass,
		LimitedTime:    membership.LimitedTime,
		LimitedContent: membership.LimitedContent,
		Description:    membership.Description,
		Details:        membership.Details,
		Price:          membership.Price,
		CreatedAt:      helpers.TimeFormat(membership.CreatedAt),
		UpdatedAt:      helpers.TimeFormat(membership.CreatedAt),
	}
}
