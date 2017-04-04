package bot

import (
	"os"

	"github.com/sad0vnikov/wundergram/bot/dialog"
	"github.com/sad0vnikov/wundergram/logger"
	"gopkg.in/telegram-bot-api.v4"
)

var dialogTreeProcessor dialog.Processor

//Init func Initializes telegram bot
func Init(token string, dialogTree dialog.Tree) {

	dialogTreeProcessor = dialog.NewProcessor(&dialogTree)

	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		panic(err)
	}

	logger.Get("main").Infof("Authorized on account %v", bot.Self.UserName)
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

		logger.Get("main").Infof("new message from %v: %v", update.Message.From.UserName, update.Message.Text)

		dialogTreeProcessor.RunNodeHandler(nextDialogNode, update.Message, bot)

	}
}

//GetTelegramBotLink returns a telegram bot t.me link
func GetTelegramBotLink() string {
	botName := os.Getenv("TELEGRAM_BOT_NAME")
	return "http://t.me/" + botName
}
