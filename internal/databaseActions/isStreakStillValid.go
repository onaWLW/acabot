package databaseActions

import "time"

func IsStreakStillValid(updateTime time.Time, messageTime time.Time) bool {
	last1312 := time.Date(messageTime.Year(), messageTime.Month(), messageTime.Day(), 13, 13, 0, 0, messageTime.Location())

	if messageTime.Before(last1312) {
		last1312 = last1312.AddDate(0, 0, -1)
	}

	return (updateTime.Day() == last1312.Day() &&
		updateTime.Month() == last1312.Month() &&
		updateTime.Year() == last1312.Year() &&
		updateTime.Hour() == 13 &&
		updateTime.Minute() == 12)
}
