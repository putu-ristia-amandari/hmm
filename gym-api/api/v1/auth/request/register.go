package request

import "gym-membership/business/user/spec"

type RegisterUserRequest struct {
	Name      string `validate:"required"`
	Email     string `validate:"required,email"`
	Password  string `validate:"required"`
	Gender    string `validate:"required"`
	City      string `validate:"required"`
	Handphone string `validate:"required"`
}

func (request *RegisterUserRequest) ToSpec() *spec.UpsertUserCreateSpec {
	return &spec.UpsertUserCreateSpec{
		Name:      request.Name,
		Email:     request.Email,
		Password:  request.Password,
		Handphone: request.Handphone,
		City:      request.City,
		Gender:    request.Gender,
	}
}
