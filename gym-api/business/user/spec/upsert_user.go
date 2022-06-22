package spec

import "time"

type UpsertUserCreateSpec struct {
	Name      string `validate:"required"`
	Email     string `validate:"required,email"`
	Password  string `validate:"required"`
	Gender    string `validate:"required"`
	City      string `validate:"required"`
	Handphone string `validate:"required"`
}

type UpsertUserUpdateSpec struct {
	Name             string `validate:"required"`
	Email            string `validate:"required,email"`
	Handphone        string `validate:"required"`
	Address          string
	City             string `validate:"required"`
	Province         string
	Nationality      string
	Gender           string `validate:"required"`
	BirthOfDate      string
	Height           int
	Weight           int
	Photo            string
	StatusMembership bool
	Status           bool
	VerifiedAt       time.Time
}

type UpsertUserUpdateProfileSpec struct {
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
	Photo       string
}

type UpsertLoginUserSpec struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

type UpsertUserPasswordUpdateSpec struct {
	OldPassword     string `validate:"required"`
	NewPassword     string `validate:"required"`
	ConfirmPassword string `validate:"required,eqfield=NewPassword"`
}

type UpsertUserForgotPasswordSpec struct {
	Email string `validate:"required,email"`
}

type UpsertUserResetPasswordSpec struct {
	NewPassword     string `validate:"required"`
	ConfirmPassword string `validate:"required,eqfield=NewPassword"`
}
