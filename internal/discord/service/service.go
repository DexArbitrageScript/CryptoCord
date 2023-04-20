package service

import (
	discord "CryptoCord/internal/discord/session"

	"github.com/bwmarrin/discordgo"
)

type App struct {
	Session *discordgo.Session
}

var DiscordService *App

func NewDiscordService() error {
	s, err := discord.NewSession()
	if err != nil {
		return err
	}
	DiscordService = &App{Session: s}
	return nil
}
