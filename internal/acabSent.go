package internal

import "github.com/bwmarrin/discordgo"

func AcabSent(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "acab" {
		s.ChannelMessageSend(m.ChannelID, "acab!")
	}
}
