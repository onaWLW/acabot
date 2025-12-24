package databaseActions

import (
	"acabot/internal/model"

	"gorm.io/gorm"
)

func AlreadySent(db *gorm.DB, uId string, sId string, hour int, minute int) bool {
	var score model.Score
	db.
		Where(&model.Score{
			UserId:   uId,
			ServerId: sId,
		}).First(&score)

	return score.LastUpdated.Hour() == hour && score.LastUpdated.Minute() == minute && score.Streak != 0
}
