package databaseActions

import (
	"time"

	"gorm.io/gorm"
)

func AlreadySent(db *gorm.DB, uId string, sId string, messageTime time.Time) bool {
	var score = GetScore(db, uId, sId)

	if score.UserId == "" {
		return false
	}

	last := score.LastUpdated

	return last.Year() == messageTime.Year() &&
		last.Month() == messageTime.Month() &&
		last.Day() == messageTime.Day() &&
		last.Hour() == 13 &&
		last.Minute() == 12 &&
		score.Streak != 0
}
