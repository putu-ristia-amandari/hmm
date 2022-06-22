package request

import "gym-membership/business/user/spec"

type ForgotPasswordRequest struct {
	Email string `json:"email"`
}

type ResetPasswordRequest struct {
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}

type UpdatePasswordRequest struct {
	OldPassword     string `json:"old_password"`
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}

func (request *ForgotPasswordRequest) ToSpec() *spec.UpsertUserForgotPasswordSpec {
	return &spec.UpsertUserForgotPasswordSpec{
		Email: request.Email,
	}
}

func (request *ResetPasswordRequest) ToSpec() *spec.UpsertUserResetPasswordSpec {
	return &spec.UpsertUserResetPasswordSpec{
		NewPassword:     request.NewPassword,
		ConfirmPassword: request.ConfirmPassword,
	}
}

func (request *UpdatePasswordRequest) ToSpec() *spec.UpsertUserPasswordUpdateSpec {
	return &spec.UpsertUserPasswordUpdateSpec{
		OldPassword:     request.OldPassword,
		NewPassword:     request.NewPassword,
		ConfirmPassword: request.ConfirmPassword,
	}
}
