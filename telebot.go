package main

import (
	"os"

	"time"

	"github.com/tucnak/telebot"
)

func main() {
	bot, err := telebot.NewBot(os.Getenv("WIKIA_TELEGRAM_API_TOKEN"))
	if err != nil {
		return
	}

	messages := make(chan telebot.Message)
	bot.Listen(messages, 1*time.Second)

	for message := range messages {
		if message.Text == "/hi" {
			bot.SendMessage(message.Chat,
				"Hello, "+message.Sender.FirstName+"!", nil)
		}
	}
}
