package response

import (
	"gym-membership/business/content"
	"gym-membership/helpers"
)

type GetContentByIDResponse struct {
	ID          int               `json:"id"`
	Category    *CategoryResponse `json:"category"`
	Title       string            `json:"title"`
	Image       string            `json:"image"`
	LinkVideo   string            `json:"link_video"`
	Description string            `json:"description"`
	CreatedBy   string            `json:"created_by"`
	UpdatedBy   string            `json:"updated_by"`
	CreatedAt   string            `json:"created_at"`
	UpdatedAt   string            `json:"updated_at"`
}

type CategoryResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	Description string `json:"description"`
}

func CreateGetAllContentResponse(contents []*content.Content) []*GetContentByIDResponse {
	var contentsResponse []*GetContentByIDResponse

	for _, content := range contents {
		contentsResponse = append(contentsResponse, CreateGetContentByIDResponse(content))
	}

	return contentsResponse
}

func CreateGetContentByIDResponse(content *content.Content) *GetContentByIDResponse {
	return &GetContentByIDResponse{
		ID: content.ID,
		Category: &CategoryResponse{
			ID:          content.Category.ID,
			Name:        content.Category.Name,
			Image:       content.Category.Image,
			Description: content.Category.Description,
		},
		Title:       content.Title,
		Image:       content.Image,
		LinkVideo:   content.LinkVideo,
		Description: content.Description,
		CreatedBy:   content.CreatedBy,
		UpdatedBy:   content.UpdatedBy,
		CreatedAt:   helpers.TimeFormat(content.CreatedAt),
		UpdatedAt:   helpers.TimeFormat(content.UpdatedAt),
	}
}
