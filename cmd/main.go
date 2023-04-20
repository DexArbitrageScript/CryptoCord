package main

import (
	"CryptoCord/bot"
	"CryptoCord/internal/discord/message"
	"CryptoCord/internal/discord/service"
	"log"
)

func main() {
	if err := service.NewDiscordService(); err != nil {
		log.Fatalf("Failed to start Discord Bot: %v", err)
	}

	s := bot.NewBotService(&bot.BotConfig{IsAllowTable: false, MsgConfig: message.MessageConfig{OnlyReceiveMsg: true}})
	s.BotOnMessage("你好", "刘总是我的大儿子！！！！哈哈哈哈")
}
