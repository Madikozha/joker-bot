import telebot
import os
from dotenv import load_dotenv
from llama_index.llms.ollama import Ollama

load_dotenv()
token=os.getenv('TG_TOKEN')

ollama = Ollama(model='llama3.2:latest', request_timeout=60.0)
bot=telebot.TeleBot(token)

@bot.message_handler(commands=['start'])
def start_message(message):
  bot.send_message(message.chat.id,"Yo ✌️ ")

@bot.message_handler(commands=['hey'])
def reply(message):
    if "rust" in message.text.lower():
        response = ollama.complete(message.text+"joke about rust proggramming language and "+"respond to this text using 2 rules: 1 use character like Joker(be funny), 2 answer shortly(2-3sentence)")
        bot.send_message(message.chat.id, response.text)
        bot.reply_to(message, "Hey, I'm rust proggramer btw 🏳️‍🌈 ")
    response = ollama.complete(message.text+"respond to this text using the following rule: use character like Joker(be funny)")
    bot.send_message(message.chat.id, response.text)
        



def main():
    bot.infinity_polling()

if __name__ == "__main__":
    main()