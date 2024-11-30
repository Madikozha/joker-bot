package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Start")
	err := godotenv.Load()
	if err != nil {
		log.Println(err.Error(), err)
	}
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_TOKEN"))
	if err != nil {
		log.Println(err.Error(), err)
	}

	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Println(err.Error(), err)
	}
	for update := range updates {
		tgChannelID := update.Message.Chat.ID
		if update.Message != nil {
			firstName := update.Message.From.FirstName
			lastName := update.Message.From.LastName
			useRespond := strings.ToLower(update.Message.Text)

			if firstName == "ToTa" && lastName == "TatO" && (strings.HasPrefix(useRespond, "hi joker") || strings.HasSuffix(useRespond, "hi joker")) {
				log.Printf("[%s] %s\n", update.Message.From.UserName, update.Message.Text)

				gif := gifHandler(tgChannelID, "https://i.imgur.com/Kd3hMX6.mp4", "Hi master!")

				bot.Send(gif)
				fmt.Println("Sending")

			} else if firstName != "ToTa" && lastName != "TatO" && (strings.HasPrefix(useRespond, "hi joker") || strings.HasSuffix(useRespond, "hi joker")) {
				log.Printf("[%s] %s\n", update.Message.From.UserName, update.Message.Text)

				gif := gifHandler(tgChannelID, "https://i.pinimg.com/originals/9f/80/73/9f807378cd83071ca8ea09e05dd03cdc.gif", "Who are you?")

				bot.Send(gif)
				fmt.Println("Sending")
			}

		}
	}
}

func main() {
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func gifHandler(tgChannelID int64, urlStr, caption string) *tgbotapi.AnimationConfig {
	gif := tgbotapi.NewAnimationShare(
		tgChannelID,
		urlStr,
	)

	gif.Caption = caption

	return &gif
}
