package internal

import (
	"acabot/internal/discordHandlers"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"
)

func HandleBot(dg *discordgo.Session, db *gorm.DB) {
	dg.AddHandler(discordHandlers.AcabSent(db))

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err := dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
}
