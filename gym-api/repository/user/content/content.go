package content

import (
	"gym-membership/business/content"
	"gym-membership/database"

	"gorm.io/gorm"
)

func ContentRepository(dbCon *database.DatabaseConnection) content.Repository {
	var contentRepository content.Repository
	contentRepository = CreateMySQlRepositoryContent(dbCon.MySQlDB)

	return contentRepository
}

type MySQLRepository struct {
	db *gorm.DB
}

func CreateMySQlRepositoryContent(db *gorm.DB) *MySQLRepository {
	return &MySQLRepository{
		db: db,
	}
}

func (repository *MySQLRepository) GetAllContent() (contents []*content.Content, err error) {
	err = repository.db.Find(&contents).Error
	if err != nil {
		return nil, err
	}

	return contents, nil
}

func (repository *MySQLRepository) GetContentByID(ID int) (content *content.Content, err error) {
	err = repository.db.First(&content, ID).Error
	if err != nil {
		return nil, err
	}

	return content, nil
}

func (repository *MySQLRepository) CreateContent(content *content.Content) (*content.Content, error) {
	err := repository.db.Model(content).Create(&content).Error
	if err != nil {
		return nil, err
	}

	return content, nil
}

func (repository *MySQLRepository) UpdateContent(content *content.Content, ID int) (*content.Content, error) {
	err := repository.db.Model(&content).Where("id=?", ID).Updates(&content).Error
	if err != nil {
		return nil, err
	}

	return content, nil
}

func (repository *MySQLRepository) DeleteContent(ID int) (content *content.Content, err error) {
	err = repository.db.First(&content, ID).Error
	if err != nil {
		return nil, err
	}

	err = repository.db.Delete(&content, ID).Error
	if err != nil {
		return nil, err
	}

	return content, nil
}
