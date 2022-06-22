package user

import (
	"gym-membership/business/membership"
	"time"
)

type User struct {
	ID               int
	MembershipID     int
	Name             string
	Email            string
	Password         string
	Handphone        string
	Address          string
	City             string
	Province         string
	Nationality      string
	Gender           string
	BirthOfDate      time.Time
	Height           int
	Weight           int
	Photo            string
	Status           bool
	StatusMembership bool
	RememberToken    string
	IsReset          bool
	VerifiedAt       time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time

	Membership membership.Membership
}

func NewUser(
	name string,
	email string,
	password string,
	gender string,
	city string,
	handphone string,
	verifyCode string,
	createdAt time.Time,
) *User {
	return &User{
		Name:          name,
		Email:         email,
		Password:      password,
		Gender:        gender,
		City:          city,
		Handphone:     handphone,
		RememberToken: verifyCode,
		CreatedAt:     createdAt,
		VerifiedAt:    time.Time{},
	}
}

func (oldUser *User) ModifyUser(
	name string,
	email string,
	handphone string,
	address string,
	city string,
	province string,
	nationality string,
	gender string,
	birthOfDate time.Time,
	height int,
	weight int,
	photo string,
	rememberToken string,
	verifiedAt time.Time,
	updatedAt time.Time,
) *User {
	return &User{
		ID:            oldUser.ID,
		Name:          name,
		Email:         email,
		Handphone:     handphone,
		Address:       address,
		City:          city,
		Province:      province,
		Nationality:   nationality,
		Gender:        gender,
		BirthOfDate:   birthOfDate,
		Height:        height,
		Weight:        weight,
		Photo:         photo,
		RememberToken: rememberToken,
		VerifiedAt:    verifiedAt,
		UpdatedAt:     updatedAt,
	}
}
