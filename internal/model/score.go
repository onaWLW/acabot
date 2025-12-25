package model

import "time"

type Score struct {
	UserId      string `gorm:"primaryKey"`
	ServerId    string `gorm:"primaryKey"`
	Streak      int
	AcabCount   int
	UserName    string
	LastUpdated time.Time
}
