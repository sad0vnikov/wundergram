package commands

import (
	"strconv"

	"github.com/sad0vnikov/wundergram/storage/daily_notifications"

	"gopkg.in/telegram-bot-api.v4"
)

func showDailyNotificationsTimeSelector(message *tgbotapi.Message, bot *tgbotapi.BotAPI) {

	buttonRows := make([][]tgbotapi.KeyboardButton, 0)

	timeButtons := [24]string{}
	for i := 0; i <= 23; i++ {
		timeButtonText := strconv.Itoa(i) + ":00"
		if i < 10 {
			timeButtonText = "0" + timeButtonText
		}
		timeButtons[i] = timeButtonText
	}

	rowWidth := 4
	for i := 0; i <= 23; i += rowWidth {
		buttons := make([]tgbotapi.KeyboardButton, 0)
		for j := 0; j < rowWidth; j++ {
			buttons = append(buttons, tgbotapi.NewKeyboardButton(timeButtons[i+j]))
		}
		buttonRows = append(buttonRows, tgbotapi.NewKeyboardButtonRow(buttons...))
	}

	buttonRows = append(buttonRows, tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Forget it. I changed my mind"),
	))

	keyboard := tgbotapi.NewReplyKeyboard(buttonRows...)

	msg := tgbotapi.NewMessage(message.Chat.ID, "When should I send you a daily notification? Choose a time a enter a more precise time, i.e. '13:25'")
	msg.ReplyMarkup = keyboard
	bot.Send(msg)

}

func enableDailyNotifications(message *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	userID := message.From.ID
	notificationTime := message.Text
	err := daily_notifications.EnableNotificationsForUser(userID, notificationTime)
	if err == nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Ok. I'll send you a daily notification at "+notificationTime)
		bot.Send(msg)
	}

}
