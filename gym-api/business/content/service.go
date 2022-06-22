package content

import (
	"gym-membership/business/content/spec"
	"gym-membership/config"

	validator "github.com/go-playground/validator/v10"
)

type Repository interface {
	GetAllContent() (content []*Content, err error)
	GetContentByID(ID int) (content *Content, err error)
	CreateContent(content *Content) (*Content, error)
	UpdateContent(contentCurrent *Content, IDCurrent int) (*Content, error)
	DeleteContent(ID int) (content *Content, err error)
}

type Service interface {
	GetAllContent() (content []*Content, err error)
	GetContentByID(ID int) (content *Content, err error)
	CreateContent(upsertContentSpec *spec.UpsertContentCreateSpec) (*Content, error)
	UpdateContent(upsertContentSpec *spec.UpsertContentUpdateSpec, IDCurrent int) (*Content, error)
	DeleteContent(ID int) (content *Content, err error)
}

type contentService struct {
	repository Repository
	config     *config.AppConfig
	validate   *validator.Validate
}

// func CreateServiceContent(repository Repository, config *config.AppConfig) Service {
// 	return &contentService{
// 		repository: repository,
// 		validate:   validator.New(),
// 		config:     config,
// 	}
// }

func (service *contentService) GetAllContent() (content []*Content, err error) {
	return service.repository.GetAllContent()
}

func (service *contentService) GetContentByID(ID int) (content *Content, err error) {
	return service.repository.GetContentByID(ID)
}

func (service *contentService) DeleteContent(ID int) (content *Content, err error) {
	return service.repository.DeleteContent(ID)
}
