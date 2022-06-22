package request

import (
	"gym-membership/business/user/spec"
	"time"
)

type UpdateMemberRequest struct {
	Name           string `validate:"required"`
	Image          string `validate:"required"`
	LimitedClass   int
	LimitedTime    int
	LimitedContent int
	Description    string
	Details        string
	Price          float64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (request *UpdateMemberRequest) ToSpec() *spec.UpsertUserUpdateSpec {
	return &spec.UpsertUserUpdateSpec{

		Name: request.Name,
		// Image:          request.Image,
		// LimitedClass:   request.LimitedClass,
		// LimitedTime:    request.LimitedTime,
		// LimitedContent: request.LimitedContent,
		// Description:    request.Description,
		// Details:        request.Details,
		// Price:          request.Price,
		// CreatedAt:      request.CreatedAt,
		// UpdatedAt:      request.UpdatedAt,
	}
}
