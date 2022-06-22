package database

import (
	"gorm.io/gorm"
)

var db *gorm.DB

func MigrateDB() error {
	return db.AutoMigrate()
}
