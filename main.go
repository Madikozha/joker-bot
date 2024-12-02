package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

func main() {
	log.Println("Start")
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err.Error())
	}
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_TOKEN"))
	if err != nil {
		log.Fatalln(err.Error())
	}

	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Println(err.Error())
	}
	for update := range updates {
		if !update.Message.IsCommand() {
			continue
		}

		tgChatID := update.Message.Chat.ID
		firstName := update.Message.From.FirstName
		lastName := update.Message.From.LastName
		// useRespond := strings.ToLower(update.Message.Text)
		gif := tgbotapi.AnimationConfig{}
		msg := tgbotapi.NewMessage(tgChatID, "")

		switch update.Message.Command() {
		case "help":
			msg.Text = "I have commands /hi and /help(this message) UwU"
			bot.Send(msg)
		case "hi":
			if firstName == "ToTa" && lastName == "TatO" {
				gif = *gifHandler(tgChatID, "https://i.imgur.com/Kd3hMX6.mp4", "Hi master!")
			} else if firstName != "ToTa" && lastName != "TatO" {
				gif = *gifHandler(tgChatID, "https://i.pinimg.com/originals/9f/80/73/9f807378cd83071ca8ea09e05dd03cdc.gif", "Who are you?")
			}
			log.Println("Sendig")
			bot.Send(gif)
		}
	}
}

func gifHandler(tgChatID int64, urlStr, caption string) *tgbotapi.AnimationConfig {
	gif := tgbotapi.NewAnimationShare(
		tgChatID,
		urlStr,
	)

	gif.Caption = caption

	return &gif
}
