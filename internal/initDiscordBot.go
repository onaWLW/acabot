package internal

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func InitDiscordBot(dg *discordgo.Session) {
	dg.AddHandler(AcabSent)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err := dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
}
