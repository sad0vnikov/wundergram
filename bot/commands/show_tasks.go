package commands

import (
	"github.com/sad0vnikov/wundergram/wunderlist/wservice"
	"gopkg.in/telegram-bot-api.v4"
)

func showTodayTasksCommand(message *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	userID := message.From.ID
	tasks, err := wservice.GetUserTodayTasks(userID)

	messageText := ""
	if err != nil {
		messageText = "Sorry! Some error occured"
		msg := tgbotapi.NewMessage(message.Chat.ID, messageText)
		bot.Send(msg)
		return
	}

	if len(tasks) > 0 {
		messageText = "Well, here are your tasks for today:\n"
		for _, t := range tasks {
			messageText += "▫ " + t.Title + "\n"
		}
	} else {
		messageText = "Hmmm, looks like you have no tasks for today"
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, messageText)
	sendMessageWithLogging(bot, msg)
}
