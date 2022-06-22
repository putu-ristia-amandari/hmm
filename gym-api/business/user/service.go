package user

import (
	"errors"
	"gym-membership/business/mail"
	"gym-membership/business/user/spec"
	"gym-membership/config"
	"gym-membership/helpers"
	"time"

	validator "github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
)

type Repository interface {
	GetAllUser() (user []*User, err error)
	GetUserByID(ID int) (user *User, err error)
	GetUserByEmail(email string) (user *User, err error)
	GetUserByVerifyCode(verifyCode string) (user *User, err error)
	CreateUser(user *User) (*User, error)
	UpdateUser(userCurrent *User, IDCurrent int) (*User, error)
	DeleteUser(ID int) (user *User, err error)
}

type Service interface {
	GetAllUser() (user []*User, err error)
	GetUserLogin(jwtToken string) (user *User, err error)
	GetUserByID(ID int) (user *User, err error)
	GetUserByEmailAndPassword(upsertUserLoginSpec *spec.UpsertLoginUserSpec) (user *User, err error)
	VerifyEmail(verifyCode string) (user *User, err error)
	ForgotPassword(upsertUserForgotPasswordSpec *spec.UpsertUserForgotPasswordSpec) (user *User, err error)
	ResetPassword(upsertUserResetPasswordSpec *spec.UpsertUserResetPasswordSpec, verifyCode string) (*User, error)
	CreateUser(upsertUserSpec *spec.UpsertUserCreateSpec) (*User, error)
	UpdateUser(upsertUserSpec *spec.UpsertUserUpdateSpec, IDCurrent int) (*User, error)
	UpdateUserPassword(upsertUserSpec *spec.UpsertUserPasswordUpdateSpec, IDCurrent int) (*User, error)
	UpdatePassword(upsertUserSpec *spec.UpsertUserResetPasswordSpec, IDCurrent int) (*User, error)
	DeleteUser(ID int) (user *User, err error)
}

type userService struct {
	repository Repository
	config     *config.AppConfig
	validate   *validator.Validate
}

func CreateService(repository Repository, config *config.AppConfig) Service {
	return &userService{
		repository: repository,
		validate:   validator.New(),
		config:     config,
	}
}

func (service *userService) GetAllUser() (user []*User, err error) {
	return service.repository.GetAllUser()
}

func (service *userService) GetUserByID(ID int) (user *User, err error) {
	return service.repository.GetUserByID(ID)
}

func (service *userService) VerifyEmail(verifyCode string) (user *User, err error) {
	user, err = service.repository.GetUserByVerifyCode(verifyCode)
	if err != nil {
		return nil, err
	}

	if user.IsReset {
		return nil, errors.New("user is reset")
	}

	user.VerifiedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Status = true
	user.RememberToken = ""

	return service.repository.UpdateUser(user, user.ID)
}

func (service *userService) ForgotPassword(upsertUserForgotPasswordSpec *spec.UpsertUserForgotPasswordSpec) (user *User, err error) {
	err = service.validate.Struct(upsertUserForgotPasswordSpec)
	if err != nil {
		return nil, err
	}

	user, err = service.repository.GetUserByEmail(upsertUserForgotPasswordSpec.Email)
	if err != nil {
		return nil, errors.New("email not found")
	}

	if user.IsReset {
		return nil, errors.New("user is reset")
	}

	time := time.Time{}
	if user.VerifiedAt == time {
		return nil, errors.New("user not verified")
	}

	if !user.Status {
		return nil, errors.New("user is inactive")
	}

	verifyCode := helpers.RandomString(64)
	user.RememberToken = verifyCode
	user.IsReset = true
	user.Status = false

	userData, err := service.repository.UpdateUser(user, user.ID)
	if err != nil {
		return nil, err
	}

	To := userData.Email
	Subject := "Reset your password"
	Body := "Reset your password by clicking this link: " + service.config.App.URL + "/api/v1/reset-password/" + user.RememberToken
	From := service.config.Mail.Username

	mailData := mail.NewMail(
		From,
		To,
		Subject,
		Body,
		"reset",
	)

	mailService := mail.CreateService(service.config)
	_, err = mailService.SendMail(mailData)
	if err != nil {
		return nil, err
	}

	return userData, nil
}

func (service *userService) ResetPassword(upsertUserResetPasswordSpec *spec.UpsertUserResetPasswordSpec, verifyCode string) (*User, error) {
	err := service.validate.Struct(upsertUserResetPasswordSpec)
	if err != nil {
		return nil, err
	}

	user, err := service.repository.GetUserByVerifyCode(verifyCode)
	if err != nil {
		return nil, err
	}

	if !user.IsReset {
		return nil, errors.New("user is not reset")
	}

	passwordHash, err := helpers.HashPassword(upsertUserResetPasswordSpec.NewPassword)
	if err != nil {
		return nil, err
	}

	user.Password = passwordHash
	user.RememberToken = ""
	user.IsReset = false
	user.Status = true
	user.UpdatedAt = time.Now()

	userData, err := service.repository.UpdateUser(user, user.ID)
	if err != nil {
		return nil, err
	}

	To := userData.Email
	Subject := "Password reset successful"
	Body := "Your password has been reset successfully, please login with your new password"
	From := service.config.Mail.Username

	mailData := mail.NewMail(
		From,
		To,
		Subject,
		Body,
		"reset-success",
	)

	mailService := mail.CreateService(service.config)
	_, err = mailService.SendMail(mailData)
	if err != nil {
		return nil, err
	}

	return userData, nil
}

func (service *userService) GetUserByEmailAndPassword(upsertUserLoginSpec *spec.UpsertLoginUserSpec) (user *User, err error) {
	err = service.validate.Struct(upsertUserLoginSpec)
	if err != nil {
		return nil, err
	}

	user, err = service.repository.GetUserByEmail(upsertUserLoginSpec.Email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	time := time.Time{}
	if user.VerifiedAt == time {
		return nil, errors.New("user not verified")
	}

	if user.IsReset {
		return nil, errors.New("user is reset")
	}

	if !user.Status {
		return nil, errors.New("user is inactive")
	}

	isTrue, err := helpers.CheckPasswordHash(upsertUserLoginSpec.Password, user.Password)
	if !isTrue || err != nil {
		return nil, errors.New("invalid password")
	}

	return user, nil
}

func (service *userService) CreateUser(upsertUserSpec *spec.UpsertUserCreateSpec) (*User, error) {
	err := service.validate.Struct(upsertUserSpec)
	if err != nil {
		return nil, err
	}

	password, err := helpers.HashPassword(upsertUserSpec.Password)
	if err != nil {
		return nil, err
	}

	verifyCode := helpers.RandomString(64)

	user := NewUser(
		upsertUserSpec.Name,
		upsertUserSpec.Email,
		password,
		upsertUserSpec.Gender,
		upsertUserSpec.City,
		upsertUserSpec.Handphone,
		verifyCode,
		time.Now(),
	)

	userData, err := service.repository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	To := userData.Email
	Subject := "Verify your email"
	Body := "Verify your email by clicking this link: " + service.config.App.URL + "/api/v1/verify/" + verifyCode
	From := service.config.Mail.Username

	mailData := mail.NewMail(
		From,
		To,
		Subject,
		Body,
		"verify",
	)

	mailService := mail.CreateService(service.config)
	_, err = mailService.SendMail(mailData)
	if err != nil {
		return nil, err
	}

	return userData, nil
}

func (service *userService) UpdateUser(upsertUserSpec *spec.UpsertUserUpdateSpec, IDCurrent int) (*User, error) {
	err := service.validate.Struct(upsertUserSpec)
	if err != nil {
		return nil, err
	}

	userCurrent, err := service.repository.GetUserByID(IDCurrent)
	if err != nil {
		return nil, err
	}

	verifyCode := userCurrent.RememberToken
	VerifiedAt := userCurrent.VerifiedAt
	if upsertUserSpec.Email != userCurrent.Email {
		verifyCode = helpers.RandomString(64)
		VerifiedAt = time.Time{}

		To := upsertUserSpec.Email
		Subject := "Verify your email"
		Body := "Verify your email by clicking this link: " + service.config.App.URL + "/api/v1/verify/" + verifyCode
		From := service.config.Mail.Username

		mailData := mail.NewMail(
			From,
			To,
			Subject,
			Body,
			"verify",
		)

		mailService := mail.CreateService(service.config)
		_, err = mailService.SendMail(mailData)
		if err != nil {
			return nil, err
		}
	}

	birthOfDate, err := time.Parse("2006-01-02", upsertUserSpec.BirthOfDate)
	if err != nil {
		return nil, err
	}

	user := userCurrent.ModifyUser(
		upsertUserSpec.Name,
		upsertUserSpec.Email,
		upsertUserSpec.Handphone,
		upsertUserSpec.Address,
		upsertUserSpec.City,
		upsertUserSpec.Province,
		upsertUserSpec.Nationality,
		upsertUserSpec.Gender,
		birthOfDate,
		upsertUserSpec.Height,
		upsertUserSpec.Weight,
		upsertUserSpec.Photo,
		verifyCode,
		VerifiedAt,
		time.Now(),
	)

	return service.repository.UpdateUser(user, IDCurrent)
}

func (service *userService) UpdateUserPassword(upsertUserSpec *spec.UpsertUserPasswordUpdateSpec, IDCurrent int) (*User, error) {
	err := service.validate.Struct(upsertUserSpec)
	if err != nil {
		return nil, err
	}

	user, err := service.repository.GetUserByID(IDCurrent)
	if err != nil {
		return nil, err
	}

	isTrue, err := helpers.CheckPasswordHash(upsertUserSpec.OldPassword, user.Password)
	if !isTrue || err != nil {
		return nil, errors.New("invalid old password")
	}

	password, err := helpers.HashPassword(upsertUserSpec.NewPassword)
	if err != nil {
		return nil, err
	}

	user.Password = password

	return service.repository.UpdateUser(user, IDCurrent)
}

func (service *userService) UpdatePassword(upsertUserSpec *spec.UpsertUserResetPasswordSpec, IDCurrent int) (*User, error) {
	err := service.validate.Struct(upsertUserSpec)
	if err != nil {
		return nil, err
	}

	user, err := service.repository.GetUserByID(IDCurrent)
	if err != nil {
		return nil, err
	}

	password, err := helpers.HashPassword(upsertUserSpec.NewPassword)
	if err != nil {
		return nil, err
	}

	user.Password = password

	return service.repository.UpdateUser(user, IDCurrent)
}

func (service *userService) GetUserLogin(jwtToken string) (user *User, err error) {
	claim := jwt.MapClaims{}
	token, _ := jwt.ParseWithClaims(jwtToken, claim, func(t *jwt.Token) (interface{}, error) {
		return []byte(service.config.App.Key), nil
	})

	jwtSignedMethod := jwt.SigningMethodHS256
	method, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok || method != jwtSignedMethod {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	var userID float64 = claims["userId"].(float64)

	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	user, err = service.repository.GetUserByID(int(userID))
	if err != nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (service *userService) DeleteUser(ID int) (user *User, err error) {
	return service.repository.DeleteUser(ID)
}
