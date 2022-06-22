package response

import (
	"gym-membership/business/instructor"
	"gym-membership/helpers"
)

type GetInstructorByIDResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Handphone   string `json:"handphone"`
	Address     string `json:"address"`
	City        string `json:"city"`
	Province    string `json:"province"`
	Nationality string `json:"nationality"`
	Gender      string `json:"gender"`
	BirthOfDate string `json:"birth_of_date"`
	Height      int    `json:"height"`
	Weight      int    `json:"Weight"`
	Photo       string `json:"photo"`
	Status      bool   `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func CreateGetAllInstructorResponse(instructors []*instructor.Instructor) []*GetInstructorByIDResponse {
	var instructorsResponse []*GetInstructorByIDResponse

	for _, instructor := range instructors {
		instructorsResponse = append(instructorsResponse, CreateGetInstructorByIDResponse(instructor))
	}

	return instructorsResponse
}

func CreateGetInstructorByIDResponse(instructor *instructor.Instructor) *GetInstructorByIDResponse {
	return &GetInstructorByIDResponse{
		ID:          instructor.ID,
		Name:        instructor.Name,
		Email:       instructor.Email,
		Handphone:   instructor.Handphone,
		Address:     instructor.Address,
		City:        instructor.City,
		Province:    instructor.Province,
		Nationality: instructor.Nationality,
		Gender:      instructor.Gender,
		BirthOfDate: helpers.DateFormat(instructor.BirthOfDate),
		Height:      instructor.Height,
		Weight:      instructor.Weight,
		Photo:       instructor.Photo,
		Status:      instructor.Status,
		CreatedAt:   helpers.TimeFormat(instructor.CreatedAt),
		UpdatedAt:   helpers.TimeFormat(instructor.UpdatedAt),
	}
}
