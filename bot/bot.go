package bot

import (
	"os"

	"github.com/sad0vnikov/wundergram/bot/dialog"
	"github.com/sad0vnikov/wundergram/logger"
	"gopkg.in/telegram-bot-api.v4"
)

var dialogTreeProcessor dialog.Processor

//Bot is a struct representing Bot state
type Bot struct {
	API *tgbotapi.BotAPI
}

//Create returns a new Bot
func Create(token string) Bot {

	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		panic(err)
	}

	return Bot{API: bot}
}

//Init func Initializes telegram bot
func (bot Bot) Init(dialogTree dialog.Tree) {

	dialogTreeProcessor = dialog.NewProcessor(&dialogTree)

	logger.Get("main").Infof("Authorized on account %#v", bot.API.Self.UserName)
	u := tgbotapi.NewUpdate(0)

	u.Timeout = 60

	updates, err := bot.API.GetUpdatesChan(u)
	if err != nil {
		panic(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		command := update.Message.Command()
		if len(command) == 0 {
			command = update.Message.Text
		}

		nextDialogNode := dialogTreeProcessor.GetNodeToMoveIn(update.Message, bot.API)

		logger.Get("main").Infof("new message from %v: %v", update.Message.From.UserName, update.Message.Text)

		go dialogTreeProcessor.RunNodeHandler(nextDialogNode, update.Message, bot.API)

	}
}

//GetTelegramBotLink returns a telegram bot t.me link
func GetTelegramBotLink() string {
	botName := os.Getenv("TELEGRAM_BOT_NAME")
	return "http://t.me/" + botName
}
