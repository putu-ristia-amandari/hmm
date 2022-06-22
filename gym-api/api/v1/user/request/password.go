package request

import "gym-membership/business/user/spec"

type UpdatePasswordRequest struct {
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}

func (request *UpdatePasswordRequest) ToSpec() *spec.UpsertUserResetPasswordSpec {
	return &spec.UpsertUserResetPasswordSpec{
		NewPassword:     request.NewPassword,
		ConfirmPassword: request.ConfirmPassword,
	}
}
