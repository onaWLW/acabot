package databaseActions

import (
	"acabot/internal/model"
	"time"

	"gorm.io/gorm"
)

func AddAcab(db *gorm.DB, uId string, sId string, uName string, messageTime time.Time) {
	var score model.Score
	db.
		Where(&model.Score{
			UserId:   uId,
			ServerId: sId,
		}).
		Attrs(model.Score{
			Streak:      0,
			AcabCount:   0,
			UserName:    uName,
			LastUpdated: messageTime,
		}).
		FirstOrCreate(&score)

	score.AcabCount += 1
	if IsStreakStillValid(score.LastUpdated, messageTime) {
		score.Streak += 1
	} else {
		score.Streak = 1
	}
	score.LastUpdated = messageTime
	score.UserName = uName

	db.Save(&score)
}
