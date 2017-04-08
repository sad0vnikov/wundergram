package commands

import (
	"gopkg.in/telegram-bot-api.v4"
)

func sendErrorResponse(chatID int64, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(chatID, "Sowwy :( An internal error occured")
	sendMessageWithLogging(bot, msg)
}
