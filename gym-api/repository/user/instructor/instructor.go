package instructor

import (
	"gym-membership/business/instructor"
	"gym-membership/database"

	"gorm.io/gorm"
)

func InstructorRepository(dbCon *database.DatabaseConnection) instructor.Repository {
	var instructorRepository instructor.Repository
	instructorRepository = CreateMySQlRepositoryInstructor(dbCon.MySQlDB)

	return instructorRepository
}

type MySQLRepository struct {
	db *gorm.DB
}

func CreateMySQlRepositoryInstructor(db *gorm.DB) *MySQLRepository {
	return &MySQLRepository{
		db: db,
	}
}

func (repository *MySQLRepository) GetAllInstructor() (instructors []*instructor.Instructor, err error) {
	err = repository.db.Find(&instructors).Error
	if err != nil {
		return nil, err
	}

	return instructors, nil
}

func (repository *MySQLRepository) GetInstructorByID(ID int) (instructor *instructor.Instructor, err error) {
	err = repository.db.First(&instructor, ID).Error
	if err != nil {
		return nil, err
	}

	return instructor, nil
}

func (repository *MySQLRepository) CreateInstructor(instructor *instructor.Instructor) (*instructor.Instructor, error) {
	err := repository.db.Model(instructor).Create(&instructor).Error
	if err != nil {
		return nil, err
	}

	return instructor, nil
}

func (repository *MySQLRepository) UpdateInstructor(instructor *instructor.Instructor, ID int) (*instructor.Instructor, error) {
	err := repository.db.Model(&instructor).Where("id=?", ID).Updates(&instructor).Error
	if err != nil {
		return nil, err
	}

	return instructor, nil
}

func (repository *MySQLRepository) DeleteInstructor(ID int) (instructor *instructor.Instructor, err error) {
	err = repository.db.First(&instructor, ID).Error
	if err != nil {
		return nil, err
	}

	err = repository.db.Delete(&instructor, ID).Error
	if err != nil {
		return nil, err
	}

	return instructor, nil
}
