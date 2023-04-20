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

	// table header
	header := []string{
		"Name",
		"Data",
		"PriceChange",
		"Exchange",
	}

	// rows data
	rows := [][]string{
		{"BNB", "2022-02-12", "86%+", "Binance"},
		{"Rndr", "2023-04-04", "25%+", "OKX"},
		{"Snx", "2023-04-19", "10%+", "Binance"},
		{"Kdm", "2023-04-18", "42%+", "Binance"},
	}

	// Initialize bot service
	/**

		  Example charts

	      TableType: <bot.Table>
		  +------+------------+-------------+----------+
		  | Name | Data       | PriceChange | Exchange |
		  +------+------------+-------------+----------+
		  | BNB  | 2022-02-12 | 86%+        | Binance  |
		  | Rndr | 2023-04-04 | 25%+        | OKX      |
		  | Snx  | 2023-04-19 | 10%+        | Binance  |
		  | Kdm  | 2023-04-18 | 42%+        | Binance  |
		  +------+------------+-------------+----------+


	      TableType: <bot.PlusTable>
		  +──────+────────────+─────────────+──────────+
		  │ Name │ Data       │ PriceChange │ Exchange │
		  +──────+────────────+─────────────+──────────+
		  │ BNB  │ 2022-02-12 │ 86%+        │ Binance  │
		  │ Rndr │ 2023-04-04 │ 25%+        │ OKX      │
		  │ Snx  │ 2023-04-19 │ 10%+        │ Binance  │
		  │ Kdm  │ 2023-04-18 │ 42%+        │ Binance  │
		  +──────+────────────+─────────────+──────────+


	      TableType: <bot.AsciiTable>
		  Name  Data        PriceChange  Exchange
		  ------ ------------ ------------- ----------
		  BNB   2022-02-12  86%+         Binance
		  Rndr  2023-04-04  25%+         OKX
		  Snx   2023-04-19  10%+         Binance
		  Kdm   2023-04-18  42%+         Binance
	*/
	s := bot.NewBotService(&bot.BotConfig{
		IsAllowTable:  true,
		TableHeader:   header,
		TableProperty: rows,
		TableType:     bot.PlusTable,
		MsgConfig: message.MessageConfig{
			OnlyReceiveMsg: true,
		}})

	s.BotOnMessageWithTable("hello")

	// Send a specific command
	// s.BotOnMessage("/about", "I am good, what about you ?")
}
