package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var bot *tgbotapi.BotAPI

func init() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_TOKEN"))
	if err != nil {
		log.Fatalf("Error initializing bot: %v", err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
}
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Start")

	var update tgbotapi.Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		log.Println("Failed to decode update:", err)
		return
	}
	if update.Message == nil {
		return
	}
	tgChatID := update.Message.Chat.ID
	firstName := update.Message.From.FirstName
	lastName := update.Message.From.LastName
	useRespond := strings.ToLower(update.Message.Text)

	if firstName == "ToTa" && lastName == "TatO" && (strings.HasPrefix(useRespond, "hi joker") || strings.HasSuffix(useRespond, "hi joker")) {
		log.Printf("[%s] %s\n", update.Message.From.UserName, update.Message.Text)

		gif := gifHandler(tgChatID, "https://i.imgur.com/Kd3hMX6.mp4", "Hi master!")

		bot.Send(gif)
		fmt.Println("Sending")

	} else if firstName != "ToTa" && lastName != "TatO" && (strings.HasPrefix(useRespond, "hi joker") || strings.HasSuffix(useRespond, "hi joker")) {
		log.Printf("[%s] %s\n", update.Message.From.UserName, update.Message.Text)

		gif := gifHandler(tgChatID, "https://i.pinimg.com/originals/9f/80/73/9f807378cd83071ca8ea09e05dd03cdc.gif", "Who are you?")

		bot.Send(gif)
		fmt.Println("Sending")

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

func main() {
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
