package discordHandlers

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func getMessageTime(m *discordgo.MessageCreate) time.Time {
	// Debug mode
	if len(m.Content) > 0 && m.Content[0] == '!' {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file: %s", err)
		}

		if os.Getenv("DEBUG") == "enabled" {
			dateStr := m.Content[1:20]
			t, err := time.Parse("2006-01-02 15:04:05", dateStr)
			if err != nil {
				log.Fatalf("Error parsing debug date: %s", err)
			}
			return t
		}
	}

	mId := m.ID
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
