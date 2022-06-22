package request

import (
	"gym-membership/business/content/spec"
)

type UpdateContentRequest struct {
	Title       string `validate:"required"`
	Image       string
	LinkVideo   string `validate:"required"`
	Description string
	CreatedBy   string `validate:"required"`
	UpdatedBy   string
}

func (request *UpdateContentRequest) ToSpec() *spec.UpsertContentUpdateSpec {
	return &spec.UpsertContentUpdateSpec{
		Title:       request.Title,
		Image:       request.Image,
		LinkVideo:   request.LinkVideo,
		Description: request.Description,
		CreatedBy:   request.CreatedBy,
		UpdatedBy:   request.UpdatedBy,
	}
}
