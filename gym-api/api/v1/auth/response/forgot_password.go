package response

type ForgotPasswordResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateForgotPasswordResponse(name string, email string) *ForgotPasswordResponse {
	return &ForgotPasswordResponse{
		Name:  name,
		Email: email,
	}
}
