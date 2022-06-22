package request

import (
	"gym-membership/business/instructor/spec"
)

type UpdateInstructorRequest struct {
	Name        string `validate:"required"`
	Handphone   string `validate:"required"`
	Address     string
	City        string `validate:"required"`
	Province    string
	Nationality string
	Gender      string `validate:"required"`
	BirthOfDate string
	Height      int
	Weight      int
	Photo       string
	Status      bool
}

func (request *UpdateInstructorRequest) ToSpec() *spec.UpsertInstructorUpdateSpec {
	return &spec.UpsertInstructorUpdateSpec{
		Name:        request.Name,
		Handphone:   request.Handphone,
		Address:     request.Address,
		City:        request.City,
		Province:    request.Province,
		Nationality: request.Nationality,
		Gender:      request.Gender,
		BirthOfDate: request.BirthOfDate,
		Height:      request.Height,
		Weight:      request.Weight,
		Photo:       request.Photo,
		Status:      request.Status,
	}
}
