package databaseActions

import (
	"acabot/internal/model"

	"gorm.io/gorm"
)

func GetScore(db *gorm.DB, uId string, sId string) model.Score {
	var score model.Score
	result := db.
		Where(&model.Score{
			UserId:   uId,
			ServerId: sId,
		}).Take(&score)

	if result.RowsAffected == 0 {
		return model.Score{}
	}
	return score
}
