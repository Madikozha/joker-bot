package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var bot *tgbotapi.BotAPI

func init() {
	var err error
	bot, err = tgbotapi.NewBotAPI(os.Getenv("TG_TOKEN"))
	if err != nil {
		log.Printf("Error initializing bot: %v", err)
		return
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Handle preflight requests
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Only handle POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var update tgbotapi.Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		log.Println("Failed to decode update:", err)
		http.Error(w, "Failed to decode update", http.StatusBadRequest)
		return
	}

	if update.Message == nil {
		w.WriteHeader(http.StatusOK)
		return
	}

	tgChatID := update.Message.Chat.ID
	firstName := update.Message.From.FirstName
	lastName := update.Message.From.LastName
	useRespond := strings.ToLower(update.Message.Text)

	var responseGif *tgbotapi.AnimationConfig
	var err error

	if firstName == "ToTa" && lastName == "TatO" &&
		(strings.HasPrefix(useRespond, "hi joker") || strings.HasSuffix(useRespond, "hi joker")) {
		log.Printf("[%s] %s\n", update.Message.From.UserName, update.Message.Text)

		responseGif = gifHandler(tgChatID, "https://i.imgur.com/Kd3hMX6.mp4", "Hi master!")

	} else if firstName != "ToTa" && lastName != "TatO" &&
		(strings.HasPrefix(useRespond, "hi joker") || strings.HasSuffix(useRespond, "hi joker")) {
		log.Printf("[%s] %s\n", update.Message.From.UserName, update.Message.Text)

		responseGif = gifHandler(tgChatID, "https://i.pinimg.com/originals/9f/80/73/9f807378cd83071ca8ea09e05dd03cdc.gif", "Who are you?")
	}

	if responseGif != nil {
		_, err = bot.Send(*responseGif)
		if err != nil {
			log.Printf("Error sending message: %v", err)
			http.Error(w, "Failed to send message", http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}

func gifHandler(tgChatID int64, urlStr, caption string) *tgbotapi.AnimationConfig {
	gif := tgbotapi.NewAnimationShare(
		tgChatID,
		urlStr,
	)

	gif.Caption = caption

	return &gif
}
