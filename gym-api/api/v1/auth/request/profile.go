package request

import "gym-membership/business/user/spec"

type UpdateProfileRequest struct {
	Name        string `validate:"required"`
	Email       string `validate:"required,email"`
	Handphone   string `validate:"required"`
	Address     string
	City        string `validate:"required"`
	Province    string
	Nationality string
	Gender      string `validate:"required"`
	BirthOfDate string
	Height      int
	Weight      int
}

func (request *UpdateProfileRequest) ToSpec() *spec.UpsertUserUpdateSpec {
	return &spec.UpsertUserUpdateSpec{
		Name:        request.Name,
		Email:       request.Email,
		Handphone:   request.Handphone,
		Address:     request.Address,
		City:        request.City,
		Province:    request.Province,
		Nationality: request.Nationality,
		Gender:      request.Gender,
		BirthOfDate: request.BirthOfDate,
		Height:      request.Height,
		Weight:      request.Weight,
	}
}
