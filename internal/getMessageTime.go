package internal

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func getMessageTime(mId string) time.Time {
	timeInt, err := strconv.ParseInt(mId, 10, 64)
	if err != nil {
		log.Fatal("Unable to convert message ID to an integer")
	}

	timeStr := fmt.Sprintf("%b", timeInt)

	for len(timeStr) != 64 {
		timeStr = "0" + timeStr
	}
	timeStr = timeStr[:42]

	timeInt, err = strconv.ParseInt(timeStr, 2, 64)
	if err != nil {
		log.Fatal("Unable to convert binary ID to an integer")
	}

	timeInt += 1420070400000

	return time.Unix(timeInt/1000, (timeInt%1000)*1_000_000)
}
