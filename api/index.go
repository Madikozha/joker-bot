package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var jokes = []string{
	"Programming is 10% writing code and 90% understanding why it’s not working",
	"Why did the math book look sad? Because it had too many problems.",
	"Knock, knock. Who’s There? Very long pause… “Java.”",
	"A guy walks into a bar and asks for 1.4 root beers. The bartender says “I’ll have to charge you extra, that’s a root beer float”. The guy says “In that case, better make it a double.",
	"There are 10 kinds of people in the world. Those who understand binary and those who don’t.",
	"I’ve got a really good UDP joke to tell you, but I don’t know if you’ll get it",
	"Things aren’t always #000000 and #FFFFFF",
	"Why do Java programmers have to wear glasses? Because they can't C",
	"Programmer: An organism that turns coffee into software",
	"Physics is the universe’s operating system",
	"There’s no place like 127.0.0.1",
	"Why do programmers take so long in the shower? They read the directions on the shampoo bottle and follow them to the letter: Lather, rinse, and repeat.",
	"A computer programmer rushes his wife to the hospital where she gives birth to their child. The doctor first hands the baby to the programmer. “Well?” his wife says impatiently. “Is it a boy, or is it a girl?” Smiling, the programmer replies, “Yes.”",
	"My love for you has no bugs",
	"What is the most used language in programming? Profanity.",
	"Real programmers count from 0",
	"My mind is like an internet browser, 19 tabs open, 3 of them are frozen, ads popping up everywhere, I have no idea where the music is coming from",
}

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received a request")

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	var update tgbotapi.Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		log.Printf("Failed to decode update: %v", err)
		return
	}

	if update.Message == nil { // ignore non-message updates
		return
	}

	// Check for the command or text to trigger the keyboard
	if update.Message.Text == "/start" {
		sendKeyboard(bot, update.Message.Chat.ID)
		return
	}

	// Check for the specific message to respond with a joke
	if strings.HasPrefix(update.Message.Text, "y") || strings.HasPrefix(update.Message.Text, "Y") {
		sendRandomJoke(bot, update.Message.Chat.ID)
		return
	}

	// Default response
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "I didn't understand that. Try saying 'yes'.\nExample: Yes or yes and etc.")
	bot.Send(msg)

	fmt.Fprintf(w, "Update processed")
}

func sendKeyboard(bot *tgbotapi.BotAPI, chatID int64) {
	buttons := []tgbotapi.KeyboardButton{
		tgbotapi.NewKeyboardButton("Yes"),
	}

	keyboard := tgbotapi.NewReplyKeyboard(buttons)
	keyboard.ResizeKeyboard = true

	msg := tgbotapi.NewMessage(chatID, "Do you want some jokes? UwU")
	msg.ReplyMarkup = keyboard

	bot.Send(msg)
}

func sendRandomJoke(bot *tgbotapi.BotAPI, chatID int64) {
	rand.Seed(time.Now().UnixNano())     // Seed the random number generator
	randomIndex := rand.Intn(len(jokes)) // Generate a random index
	joke := jokes[randomIndex]           // Select a random joke

	msg := tgbotapi.NewMessage(chatID, joke)
	bot.Send(msg)
}

func main() {
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
