package internal

import (
	"acabot/internal/model"

	"gorm.io/gorm"
)

func InitDatabase(db *gorm.DB) {
	db.AutoMigrate(&model.Score{})
}
