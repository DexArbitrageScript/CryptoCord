package message

import (
	"CryptoCord/internal/discord/service"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

type MessageBody struct {
	App       *service.App
	Msgconfig *MessageConfig
}

type MessageConfig struct {
	// set maximum length of message
	MaxLength int
	// set prefix for the message
	Prefix string
	// set suffix for the message
	Suffix string
	// set mentions for specific users or roles
	Mentions []string

	// only care about receiving message events
	OnlyReceiveMsg bool
}

type MessageHandler = func(s *discordgo.Session, m *discordgo.MessageCreate)

type MessageService interface {
	NewMessage(app *service.App, config *MessageConfig) *MessageBody
	OnMessage(handler MessageHandler)
	SendMessage(cid string, content string)
	SendMessageWithUser(cid string, mid string, content string, msgReference *discordgo.MessageReference)
}

func (msg *MessageBody) NewMessage(app *service.App, config *MessageConfig) *MessageBody {

	// initialize MessageBody data
	return &MessageBody{
		App:       app,
		Msgconfig: config,
	}
}

func (msg *MessageBody) OnMessage(handler MessageHandler) {
	// register message handler
	msg.App.Session.AddHandler(handler)
	err := msg.App.Session.Open()
	if err != nil {
		log.Fatalf("Failed to open Discord session: %v", err)
	}
	msg.keep()

	// clean up session
	msg.cleanup()
}

func (msg *MessageBody) keep() {
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

func (msg *MessageBody) cleanup() {
	msg.App.Session.Close()
}

func (msg *MessageBody) SendMessage(cid string, content string) {
	_, err := msg.App.Session.ChannelMessageSend(cid, content)
	if err != nil {
		panic(err)
	}
}

func (msg *MessageBody) SendMessageWithUser(cid string, mid string, content string, msgReference *discordgo.MessageReference) {
	_, err := msg.App.Session.ChannelMessageSendReply(cid, content, msgReference)

	if err != nil {
		log.Fatal("Error sending reply:", err)
		return
	}

}
