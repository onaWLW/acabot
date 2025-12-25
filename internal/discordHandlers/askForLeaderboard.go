package discordHandlers

import (
	"acabot/internal/databaseActions"

	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"
)

func AskForLeaderboard(db *gorm.DB) func(*discordgo.Session, *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.ApplicationCommandData().Name == "acableaderboard" {
			leaderboard := databaseActions.GetLeaderboard(db, i.GuildID)

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: leaderboard,
				},
			})
		}
	}
}
