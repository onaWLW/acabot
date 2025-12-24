package databaseActions

import (
	"acabot/internal/model"
	"time"

	"gorm.io/gorm"
)

func AlreadySent(db *gorm.DB, uId string, sId string, messageTime time.Time) bool {
	var score model.Score
	result := db.
		Where(&model.Score{
			UserId:   uId,
			ServerId: sId,
		}).Take(&score)

	if result.RowsAffected == 0 {
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
