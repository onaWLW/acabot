package internal

import (
	"time"

	"gorm.io/gorm"
)

type Score struct {
	UserId      string `gorm:"primaryKey"`
	ServerId    string `gorm:"primaryKey"`
	Streak      int
	AcabCount   int
	AcacCount   int
	LastUpdated time.Time
}

func InitDatabase(db *gorm.DB) {
	db.AutoMigrate(&Score{})

	db.Create(&Score{UserId: "UID", ServerId: "SID", Streak: 0, AcabCount: 0, AcacCount: 0, LastUpdated: time.Now()})
}
