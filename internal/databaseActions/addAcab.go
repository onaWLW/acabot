package databaseActions

import (
	"acabot/internal/model"
	"time"

	"gorm.io/gorm"
)

func AddAcab(db *gorm.DB, uId string, sId string) {
	var score model.Score
	db.
		Where(&model.Score{
			UserId:   uId,
			ServerId: sId,
		}).
		Attrs(model.Score{
			Streak:      0,
			AcabCount:   0,
			AcacCount:   0,
			LastUpdated: time.Now(),
		}).
		FirstOrCreate(&score)

	score.AcabCount += 1
	if IsStreakStillValid(score.LastUpdated) {
		score.Streak += 1
	} else {
		score.Streak = 1
	}
	score.LastUpdated = time.Now()

	db.Save(&score)
}
