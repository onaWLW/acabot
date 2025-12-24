package discordHandlers

import (
	"acabot/internal/databaseActions"
	"log"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"
)

func AcabSent(db *gorm.DB) func(*discordgo.Session, *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		if isAcab(m.Content) && is13h12(m.ID) {
			err := s.MessageReactionAdd(m.ChannelID, m.ID, "🔥")
			if err != nil {
				log.Fatal("Unable to add a reaction to the 'acab' message")
			}

			databaseActions.AddAcab(db, "user", "server")
		}
	}
}

func isAcab(message string) bool {
	re := regexp.MustCompile(`\s`)
	message = re.ReplaceAllString(message, ``)
	message = strings.ToLower(message)

	return strings.Contains(message, "acab") || strings.Contains(message, "1312")
}

func is13h12(mId string) bool {
	date := getMessageTime(mId)
	return true
	return date.Hour() == 13 && date.Minute() == 12
}
