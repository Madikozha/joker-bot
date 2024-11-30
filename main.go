package main

import (
	"fmt"
	"os"

	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Start")
	err := godotenv.Load()
	if err != nil {
		slog.Error(err.Error(), err)
	}
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_TOKEN"))
	if err != nil {
		slog.Error(err.Error(), err)
	}

	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		slog.Error(err.Error(), err)
	}
	for update := range updates {
		if update.Message != nil {
			slog.Info("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "I AM REPLYING 2.0")
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}
	}
}
