# import telegram
# from llama_index.llms.ollama import Ollama

# # Замените на ваш токен бота и путь к модели Ollama
# bot = telegram.Bot(token='7740793577:AAGy8KBAwIW4J3KGfU-Ax_M-zdc7jKf-ckw')
# ollama = Ollama(model='llama3.2:latest')

# def handle_message(update, context):
#     text = update.message.text
#     response = ollama.generate(text)
#     bot.send_message(chat_id=update.effective_chat.id, text=response)

# updater = telegram.Update(token='7740793577:AAGy8KBAwIW4J3KGfU-Ax_M-zdc7jKf-ckw', use_context=True)
# dispatcher = updater.dispatcher
# dispatcher.add_handler(telegram.MessageHandler(telegram.Filters.text & ~telegram.Filters.command, handle_message))
# updater.start_polling()
from llama_index.llms.ollama import Ollama

ollama = Ollama(model='llama3.2:latest')
text=input()
response = ollama.complete(text+"use noSql to answer")
print(response)