package migrations

import (
	"fund-aplly-back/models"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	db.AutoMigrate(&models.Post{}, &models.User{})
}
