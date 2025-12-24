package discordHandlers

import (
	"acabot/internal/databaseActions"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"
)

func AcabSent(db *gorm.DB) func(*discordgo.Session, *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		messageTime := getMessageTime(m)

		if isAcab(m.Content) && is13h12(messageTime) && !databaseActions.AlreadySent(db, m.Author.ID, m.GuildID, messageTime) {
			err := s.MessageReactionAdd(m.ChannelID, m.ID, "🔥")
			if err != nil {
				log.Fatal("Unable to add a reaction to the 'acab' message")
			}

			databaseActions.AddAcab(db, m.Author.ID, m.GuildID, messageTime)
		}
	}
}

func isAcab(message string) bool {
	re := regexp.MustCompile(`\s`)
	message = re.ReplaceAllString(message, ``)
	message = strings.ToLower(message)

	return strings.Contains(message, "acab") || strings.Contains(message, "1312")
}

func is13h12(messageTime time.Time) bool {
	return messageTime.Hour() == 13 && messageTime.Minute() == 12
}
