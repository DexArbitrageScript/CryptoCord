package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

type App struct {
	session *discordgo.Session
}

var DiscordApp *App

func main() {
	s, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		fmt.Println("Failed to create Discord session:", err)
		return
	}

	DiscordApp = &App{session: s}
}
