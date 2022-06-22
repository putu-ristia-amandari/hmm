package membership

import (
	"gym-membership/business/membership/spec"
	"gym-membership/config"

	validator "github.com/go-playground/validator/v10"
)

type Repository interface {
	GetAllMember() (memberships []*Membership, err error)
	GetMemberByID(ID int) (membership *Membership, err error)
	CreateMember(membership *Membership) (*Membership, error)
	UpdateMember(membershipCurrent *Membership, IDCurrent int) (*Membership, error)
	DeleteMember(ID int) (membership *Membership, err error)
}

type Service interface {
	GetAllMember() (memberships []*Membership, err error)
	GetMemberByID(ID int) (membership *Membership, err error)
	CreateMember(upsertMemberSpec *spec.UpsertMemberCreateSpec) (*Membership, error)
	UpdateMember(upsertMemberSpec *spec.UpsertMemberUpdateSpec, IDCurrent int) (*Membership, error)
	DeleteMember(ID int) (membership *Membership, err error)
}

type membershipService struct {
	repository Repository
	config     *config.AppConfig
	validate   *validator.Validate
}

func CreateServiceMembership(repository Repository, config *config.AppConfig) Service {
	return &membershipService{
		repository: repository,
		validate:   validator.New(),
		config:     config,
	}
}

func (service *membershipService) GetAllMember() (membership []*Membership, err error) {
	return service.repository.GetAllMember()
}

func (service *membershipService) GetMemberByID(ID int) (membership *Membership, err error) {
	return service.repository.GetMemberByID(ID)
}

func (service *membershipService) DeleteMember(ID int) (membership *Membership, err error) {
	return service.repository.DeleteMember(ID)
}
