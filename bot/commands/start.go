package commands

import (
	"fmt"

	"github.com/sad0vnikov/wundergram/storage/timezones"
	"github.com/sad0vnikov/wundergram/util"
	"github.com/sad0vnikov/wundergram/wunderlist"
	"gopkg.in/telegram-bot-api.v4"
)

//start is an initial bot command handler
func start(message *tgbotapi.Message, bot *tgbotapi.BotAPI) {

	if message.Location != nil {
		saveUserLocation(message, bot)
	}

	isAuthorized, err := wunderlist.IsUserAuthorized(message.From.ID)
	if err != nil {
		sendErrorResponse(message.Chat.ID, bot)
		return
	}

	hasTimezoneSettings := timezones.UserHasLocation(message.From.ID)
	if isAuthorized && !hasTimezoneSettings {
		sendNeedLocationMessage(message, bot)
		return
	}

	if isAuthorized && hasTimezoneSettings {
		sendAuthorizedMessage(message, bot)
		return
	}

	arg := message.CommandArguments()

	if len(arg) > 0 {
		startDoAuth(message, bot, arg)
	} else {
		startRespondNeedAuth(message, bot)
	}
}

func sendAuthorizedMessage(message *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	msgText := "You were successfully authenticated. Now I can send you Wunderlist reminders"
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Send me daily notifications!")),
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Show me my tasks for today")),
	)

	msg := tgbotapi.NewMessage(message.Chat.ID, msgText)
	msg.ReplyMarkup = keyboard

	sendMessageWithLogging(bot, msg)

}

func sendNeedLocationMessage(message *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(
		message.Chat.ID,
		"I also need to know your location to detect your timezone. I can't send you notifications just in time without knowing your timezone",
	)
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButtonLocation("📍 Send location"),
		),
	)

	sendMessageWithLogging(bot, msg)
}

func saveUserLocation(message *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	userTimezone := util.GetTimezoneByCoord(message.Location.Latitude, message.Location.Longitude)
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Thank you! Your timezone is %v", userTimezone.String()))
	timezones.Put(message.From.ID, userTimezone)
	sendMessageWithLogging(bot, msg)

}

func startRespondNeedAuth(message *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	msgText := fmt.Sprintf(
		"Hi! You should give me access to your account so that I could send you the reminders. Follow this link to do that: %v",
		wunderlist.GetUserAuthLink(),
	)
	msg := tgbotapi.NewMessage(message.Chat.ID, msgText)
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(false)

	sendMessageWithLogging(bot, msg)

}

func startDoAuth(message *tgbotapi.Message, bot *tgbotapi.BotAPI, code string) {
	err := wunderlist.AuthorizeUser(message.From.ID, code)

	if err != nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Sorry! Some error has occured")
		bot.Send(msg)
		return
	}

	sendAuthorizedMessage(message, bot)
}
