package bot

import (
	"log"
	"os"

	"github.com/sad0vnikov/wundergram/bot/dialog"
	"gopkg.in/telegram-bot-api.v4"
)

var dialogTreeProcessor dialog.TreeProcessor

//Init func Initializes telegram bot
func Init(token string, dialogTree dialog.Tree) {

	dialogTreeProcessor = dialog.NewTreeProcessor(&dialogTree)

	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		panic(err)
	}

	log.Printf("Authorized on account %v", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)

	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		command := update.Message.Command()
		if len(command) == 0 {
			command = update.Message.Text
		}

		nextDialogNode := dialogTreeProcessor.GetNodeToMoveIn(update.Message, bot)

		log.Printf("new message from %v: %v", update.Message.From.UserName, update.Message.Text)
		nextDialogNode.Handler(update.Message, bot)
	}
}

//GetTelegramBotLink returns a telegram bot t.me link
func GetTelegramBotLink() string {
	botName := os.Getenv("TELEGRAM_BOT_NAME")
	return "http://t.me/" + botName
}
