package commands

import "gopkg.in/telegram-bot-api.v4"

func enableDailyNotifications(message *tgbotapi.Message, bot *tgbotapi.BotAPI) {

	buttonRows := make([]tgbotapi.KeyboardButton, 5)

	keyboard := tgbotapi.NewReplyKeyboard(buttonRows)

	msg := tgbotapi.NewMessage(message.Chat.ID, "When should I send you a daily notification?")
	msg.ReplyMarkup = keyboard
	bot.Send(msg)

}
