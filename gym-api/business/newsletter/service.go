package newsletter

import (
	"gym-membership/business/newsletter/spec"
	"gym-membership/config"

	validator "github.com/go-playground/validator/v10"
)

type Repository interface {
	GetAllNews() (newsletter []*NewsLetter, err error)
	GetNewsByID(ID int) (newsletter *NewsLetter, err error)
	CreateNews(newsletter *NewsLetter) (*NewsLetter, error)
	UpdateNews(newsletterCurrent *NewsLetter, IDCurrent int) (*NewsLetter, error)
	DeleteNews(ID int) (newsletter *NewsLetter, err error)
}

type Service interface {
	GetAllNews() (newsletter []*NewsLetter, err error)
	GetNewsByID(ID int) (newsletter *NewsLetter, err error)
	CreateNews(upsertNewsLetterSpec *spec.UpsertNewsLetterCreateSpec) (*NewsLetter, error)
	UpdateNews(upsertNewsLetterSpec *spec.UpsertNewsLetterUpdateSpec, IDCurrent int) (*NewsLetter, error)
	DeleteNews(ID int) (newsletter *NewsLetter, err error)
}

type newsService struct {
	repository Repository
	config     *config.AppConfig
	validate   *validator.Validate
}

// func CreateServiceNews(repository Repository, config *config.AppConfig) Service {
// 	return &newsService{
// 		repository: repository,
// 		validate:   validator.New(),
// 		config:     config,
// 	}
// }

func (service *newsService) GetAllNews() (newsletter []*NewsLetter, err error) {
	return service.repository.GetAllNews()
}

func (service *newsService) GetNewsByID(ID int) (newsletter *NewsLetter, err error) {
	return service.repository.GetNewsByID(ID)
}

func (service *newsService) DeleteNews(ID int) (newsletter *NewsLetter, err error) {
	return service.repository.DeleteNews(ID)
}
