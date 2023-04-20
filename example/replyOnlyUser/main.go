package main

import (
	"CryptoCord/bot"
	"CryptoCord/internal/discord/service"
	"log"
)

func main() {
	// create a discord service
	if err := service.NewDiscordService(); err != nil {
		log.Fatalf("Failed to start Discord Bot: %v", err)
	}

	// Initialize bot service
	cryptoCordBot := bot.NewBotService(&bot.BotConfig{})
	// send message
	cryptoCordBot.BotOnMessageWithOnlyUser("hello", "Hi, my name is CryptoCord Bot, what is your name ?")
}
