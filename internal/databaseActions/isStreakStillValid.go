package databaseActions

import "time"

func IsStreakStillValid(t time.Time) bool {
	now := time.Now()

	last1312 := time.Date(now.Year(), now.Month(), now.Day(), 13, 12, 0, 0, now.Location())

	if now.Before(last1312) {
		last1312 = last1312.AddDate(0, 0, -1)
	}

	return t.After(last1312)
}
