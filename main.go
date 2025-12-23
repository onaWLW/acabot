package main

import (
	"acabot/internal"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	db, err := gorm.Open(sqlite.Open(os.Getenv("SQLITE_DATABASE_PATH")), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	internal.InitDatabase(db)

	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Fatalf("Wrong API key: %s", err)
	}

	internal.InitDiscordBot(dg)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()
}
