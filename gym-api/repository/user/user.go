package user

import (
	"errors"
	"gym-membership/business/user"
	"gym-membership/database"

	"gorm.io/gorm"
)

func UserRepository(dbCon *database.DatabaseConnection) user.Repository {
	var userRepository user.Repository
	userRepository = CreateMySQlRepository(dbCon.MySQlDB)

	return userRepository
}

type MySQLRepository struct {
	db *gorm.DB
}

func CreateMySQlRepository(db *gorm.DB) *MySQLRepository {
	return &MySQLRepository{
		db: db,
	}
}

func (repository *MySQLRepository) GetAllUser() (users []*user.User, err error) {
	err = repository.db.Preload("Membership").Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (repository *MySQLRepository) GetUserByID(ID int) (user *user.User, err error) {
	err = repository.db.Preload("Membership").First(&user, ID).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repository *MySQLRepository) GetUserByEmail(email string) (user *user.User, err error) {
	err = repository.db.Preload("Membership").Where("email=?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repository *MySQLRepository) GetUserByVerifyCode(verifyCode string) (user *user.User, err error) {
	err = repository.db.Preload("Membership").Where("remember_token=?", verifyCode).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repository *MySQLRepository) CreateUser(user *user.User) (*user.User, error) {
	userCheck, err := repository.GetUserByEmail(user.Email)
	if err == nil && userCheck.Email == user.Email {
		return nil, errors.New("account already taken")
	}

	createUser := map[string]interface{}{
		"ID":            nil,
		"MembershipID":  nil,
		"Name":          user.Name,
		"Email":         user.Email,
		"Password":      user.Password,
		"Handphone":     user.Handphone,
		"City":          user.City,
		"Gender":        user.Gender,
		"RememberToken": user.RememberToken,
		"CreatedAt":     user.CreatedAt,
	}

	err = repository.db.Model(user).Create(&createUser).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repository *MySQLRepository) UpdateUser(userCurrent *user.User, IDCurrent int) (*user.User, error) {
	membershipID := interface{}(nil)
	membershipID = userCurrent.MembershipID
	if userCurrent.MembershipID == 0 {
		membershipID = nil
	}

	updateUser := map[string]interface{}{
		"ID":               IDCurrent,
		"MembershipID":     membershipID,
		"Name":             userCurrent.Name,
		"Email":            userCurrent.Email,
		"Password":         userCurrent.Password,
		"Handphone":        userCurrent.Handphone,
		"Address":          userCurrent.Address,
		"City":             userCurrent.City,
		"Province":         userCurrent.Province,
		"Nationality":      userCurrent.Nationality,
		"Gender":           userCurrent.Gender,
		"BirthOfDate":      userCurrent.BirthOfDate,
		"Height":           userCurrent.Height,
		"Weight":           userCurrent.Weight,
		"Photo":            userCurrent.Photo,
		"Status":           userCurrent.Status,
		"StatusMembership": userCurrent.StatusMembership,
		"RememberToken":    userCurrent.RememberToken,
		"IsReset":          userCurrent.IsReset,
		"VerifiedAt":       userCurrent.VerifiedAt,
		"CreatedAt":        userCurrent.CreatedAt,
		"UpdatedAt":        userCurrent.UpdatedAt,
	}

	err := repository.db.Model(&userCurrent).Where("id=?", IDCurrent).Updates(&updateUser).Error
	if err != nil {
		return nil, err
	}

	return userCurrent, nil
}

func (repository *MySQLRepository) DeleteUser(ID int) (user *user.User, err error) {
	err = repository.db.Preload("Membership").First(&user, ID).Error
	if err != nil {
		return nil, err
	}

	err = repository.db.Delete(&user, ID).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
