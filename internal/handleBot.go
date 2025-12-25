package internal

import (
	"acabot/internal/discordHandlers"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"
)

func HandleBot(dg *discordgo.Session, db *gorm.DB) {
	dg.AddHandler(discordHandlers.AcabSent(db))

	dg.AddHandler(discordHandlers.AskForLeaderboard(db))

	err := dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	_, err = dg.ApplicationCommandCreate(
		dg.State.User.ID,
		"968590769361387600",
		&discordgo.ApplicationCommand{
			Name:        "acableaderboard",
			Description: "Montre le leaderboard d'ACABot",
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
}
