package bot

import (
	"CryptoCord/internal/discord/message"
	"CryptoCord/internal/discord/service"
	"CryptoCord/internal/utils"
	"context"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const (
	Table      = "Table"
	PlusTable  = "PlusTable"
	AsciiTable = "AsciiTable"
)

type BotService struct {
	App        *service.App
	MessageSvc message.MessageService
	BotConfig  *BotConfig

	Event *EventMessage
}

type BotConfig struct {
	MsgConfig message.MessageConfig

	IsAllowTable  bool
	TableHeader   []string
	TableProperty [][]string
	TableType     string
}

func NewBotService(botConfig *BotConfig) *BotService {
	var messageSvc message.MessageService = &message.MessageBody{
		App:       service.DiscordService,
		Msgconfig: &botConfig.MsgConfig,
	}

	messageSvc.NewMessage(service.DiscordService, &message.MessageConfig{
		OnlyReceiveMsg: botConfig.MsgConfig.OnlyReceiveMsg})

	return &BotService{
		App:        service.DiscordService,
		MessageSvc: messageSvc,
		BotConfig:  botConfig,
	}
}

func (svc *BotService) BotOnMessage(text string, content string) {
	svc.MessageSvc.OnMessage(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Content == text {
			svc.BotSendMessage(m.ChannelID, content)
		}
	})
}

func (svc *BotService) BotOnMessageWithTable(text string) {
	svc.MessageSvc.OnMessage(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		// If IsAllowTable is set to false, return an error indicating that table generation is not allowed.
		if !svc.BotConfig.IsAllowTable {
			log.Fatal("The bot config isAllowTable is false")
			return
		}

		if m.Content == text {
			svc.BotSendMessage(m.ChannelID, " ")
		}
	})
}

func (svc *BotService) BotOnMessageWithOnlyUser(text string, content string) {
	svc.MessageSvc.OnMessage(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Content == text {
			svc.BotSendMessageOnlyUser(m.ChannelID, m.ID, content, m.Reference())
		}
	})
}

func (svc *BotService) BotOnMessageWithCommand(command string, content string) {
	svc.MessageSvc.OnMessage(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if strings.HasPrefix(m.Content, command) {
			svc.BotSendMessage(m.ChannelID, content)
		}
	})
}

func (svc *BotService) BotOnMessageWithChannel(cid string, command string, content string) {
	svc.MessageSvc.OnMessage(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Content == command && m.ChannelID == cid {
			svc.BotSendMessage(m.ChannelID, content)
		}
	})
}

func (svc *BotService) BotOnAutomaticMessage(content string) {
	//prefixedContent := fmt.Sprintf("[%s] %s", svc.CryptoStore.GetSymbol(), content)
	ctx, _ := context.WithCancel(context.Background())
	event := NewEvent(ctx)
	svc.MessageSvc.OnMessage(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		go event.RunAutomaticEvent(func() { svc.MessageSvc.SendMessage(m.ChannelID, content) })
	})
}

func (svc *BotService) BotSendMessage(cid string, content string) {
	svc.MessageSvc.SendMessage(cid, content)
}

func (svc *BotService) BotSendMessageOnlyUser(cid string, mid string, content string, msgReference *discordgo.MessageReference) {
	svc.MessageSvc.SendMessageWithUser(cid, mid, content, msgReference)
}

func (svc *BotService) handlerContent(content string) string {
	var tableContent string
	switch svc.BotConfig.TableType {
	case Table:
		tableContent = utils.GenerateTable(svc.BotConfig.TableHeader, svc.BotConfig.TableProperty)
	case PlusTable:
		tableContent = utils.GeneratePlusTable(svc.BotConfig.TableHeader, svc.BotConfig.TableProperty)
	case AsciiTable:
		tableContent = utils.GenerateAsciiTable(svc.BotConfig.TableHeader, svc.BotConfig.TableProperty)

	default:
		log.Fatal("The table type is not found")
	}

	return tableContent
}
