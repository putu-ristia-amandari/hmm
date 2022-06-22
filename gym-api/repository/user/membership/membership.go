package membership

import (
	"gym-membership/business/membership"
	"gym-membership/database"

	"gorm.io/gorm"
)

func MemberRepository(dbCon *database.DatabaseConnection) membership.Repository {
	var membershipRepository membership.Repository
	membershipRepository = CreateMySQlRepositoryMember(dbCon.MySQlDB)

	return membershipRepository
}

type MySQLRepository struct {
	db *gorm.DB
}

func CreateMySQlRepositoryMember(db *gorm.DB) *MySQLRepository {
	return &MySQLRepository{
		db: db,
	}
}

func (repository *MySQLRepository) GetAllMember() (memberships []*membership.Membership, err error) {
	err = repository.db.Find(&memberships).Error
	if err != nil {
		return nil, err
	}

	return memberships, nil
}

func (repository *MySQLRepository) GetMemberByID(ID int) (membership *membership.Membership, err error) {
	err = repository.db.First(&membership, ID).Error
	if err != nil {
		return nil, err
	}

	return membership, nil
}

func (repository *MySQLRepository) CreateMember(membership *membership.Membership) (*membership.Membership, error) {
	err := repository.db.Model(membership).Create(&membership).Error
	if err != nil {
		return nil, err
	}

	return membership, nil
}

func (repository *MySQLRepository) UpdateMember(membership *membership.Membership, ID int) (*membership.Membership, error) {
	err := repository.db.Model(&membership).Where("id=?", ID).Updates(&membership).Error
	if err != nil {
		return nil, err
	}

	return membership, nil
}

func (repository *MySQLRepository) DeleteMember(ID int) (membership *membership.Membership, err error) {
	err = repository.db.First(&membership, ID).Error
	if err != nil {
		return nil, err
	}

	err = repository.db.Delete(&membership, ID).Error
	if err != nil {
		return nil, err
	}

	return membership, nil
}
