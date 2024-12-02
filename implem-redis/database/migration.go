package database

import (
	"golang-chapter-41/implem-redis/model"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.User{},
		&model.Shipping{},
	)
}
