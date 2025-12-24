package discordHandlers

import (
	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"
)

func AskForLeaderboard(db *gorm.DB) func(*discordgo.Session, *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.ApplicationCommandData().Name == "acableaderboard" {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Pong 🏓",
				},
			})
		}
	}
}
