package instructor

import (
	"gym-membership/business/instructor/spec"
	"gym-membership/config"

	validator "github.com/go-playground/validator/v10"
)

type Repository interface {
	GetAllInstructor() (instructor []*Instructor, err error)
	GetInstructorByID(ID int) (instructor *Instructor, err error)
	CreateInstructor(instructor *Instructor) (*Instructor, error)
	UpdateInstructor(instructorCurrent *Instructor, IDCurrent int) (*Instructor, error)
	DeleteInstructor(ID int) (instructor *Instructor, err error)
}

type Service interface {
	GetAllInstructor() (instructor []*Instructor, err error)
	GetInstructorByID(ID int) (user *Instructor, err error)
	CreateInstructor(upsertInstructorSpec *spec.UpsertInstructorCreateSpec) (*Instructor, error)
	UpdateInstructor(upsertInstructorSpec *spec.UpsertInstructorUpdateSpec, IDCurrent int) (*Instructor, error)
	DeleteInstructor(ID int) (instructor *Instructor, err error)
}

type instructorService struct {
	repository Repository
	config     *config.AppConfig
	validate   *validator.Validate
}

// func CreateServiceInstructor(repository Repository, config *config.AppConfig) Service {
// 	return &instructorService{
// 		repository: repository,
// 		validate:   validator.New(),
// 		config:     config,
// 	}
// }

func (service *instructorService) GetAllInstructor() (instructor []*Instructor, err error) {
	return service.repository.GetAllInstructor()
}

func (service *instructorService) GetInstructorByID(ID int) (instructor *Instructor, err error) {
	return service.repository.GetInstructorByID(ID)
}

func (service *instructorService) DeleteInstructor(ID int) (instructor *Instructor, err error) {
	return service.repository.DeleteInstructor(ID)
}
