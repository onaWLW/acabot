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
	var streakIcons = [11]string{"", "", "2️⃣", "3️⃣", "4️⃣", "5️⃣", "6️⃣", "7️⃣", "8️⃣", "9️⃣", "🔟"}

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

			databaseActions.AddAcab(db, m.Author.ID, m.GuildID, m.Author.DisplayName(), messageTime)

			var streak = databaseActions.GetScore(db, m.Author.ID, m.GuildID).Streak

			if streak > 1 && streak <= 10 {
				err = s.MessageReactionAdd(m.ChannelID, m.ID, streakIcons[streak])
				if err != nil {
					log.Fatal("Unable to add a reaction to the 'acab' message")
				}
			} else if streak > 10 {
				err = s.MessageReactionAdd(m.ChannelID, m.ID, "💥")
				if err != nil {
					log.Fatal("Unable to add a reaction to the 'acab' message")
				}

				err = s.MessageReactionAdd(m.ChannelID, m.ID, "🧨")
				if err != nil {
					log.Fatal("Unable to add a reaction to the 'acab' message")
				}
			}
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
