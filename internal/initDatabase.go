package internal

import (
	"gorm.io/gorm"
)

type Score struct {
	gorm.Model
	UserId    string
	ServerId  string
	Streak    int
	AcabCount int
	AcacCount int
}

func InitDatabase(db *gorm.DB) {
	db.AutoMigrate(&Score{})

	db.Create(&Score{UserId: "UID", ServerId: "SID", Streak: 0, AcabCount: 0, AcacCount: 0})
}
