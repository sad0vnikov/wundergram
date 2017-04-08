package commands

import (
	"github.com/sad0vnikov/wundergram/logger"
	"gopkg.in/telegram-bot-api.v4"
)

func sendMessageWithLogging(bot *tgbotapi.BotAPI, msg tgbotapi.Chattable) {
	_, err := bot.Send(msg)
	if err != nil {
		logger.Get("main").Errorf("error sending message: %v", err)
	}

	logger.Get("main").Debug("sent message: %v", msg)
}
