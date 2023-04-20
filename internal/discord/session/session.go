package discord

import (
	"os"

	"github.com/bwmarrin/discordgo"
)

func NewSession() (*discordgo.Session, error) {
	return discordgo.New("Bot " + os.Getenv("TOKEN"))
}
