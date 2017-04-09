package tasks

import (
	"time"

	"github.com/sad0vnikov/wundergram/bot"
	"github.com/sad0vnikov/wundergram/logger"
	"github.com/sad0vnikov/wundergram/storage/daily_notifications"
	"github.com/sad0vnikov/wundergram/storage/timezones"
	"github.com/sad0vnikov/wundergram/wunderlist/wservice"
	"gopkg.in/telegram-bot-api.v4"
)

//RunDailyNotifications runs daily notifications loop
func RunDailyNotifications(bot bot.Bot) {
	checkDailyNotifications(bot.API)

	ticker := time.NewTicker(time.Minute)
	for range ticker.C {
		checkDailyNotifications(bot.API)
	}

}

func checkDailyNotifications(bot *tgbotapi.BotAPI) {
	logger.Get("main").Debug("checking for notifications to send")
	notifications, err := daily_notifications.GetAll()
	if err != nil {
		logger.Get("main").Errorf("error getting notifications list from db %+v", err)
		return
	}

	logger.Get("main").Debugf("found %v notification configs: ", len(notifications), notifications)

	for _, n := range notifications {
		userLocation, err := timezones.Get(n.UserID)
		if err != nil || userLocation == nil {
			logger.Get("main").Noticef("error getting timezone for user %v, assuming UTC. error: %v", n.UserID, err)
			userLocation = time.UTC
		}

		curTime := time.Now().In(userLocation)
		if n.CheckIsTimeToSend(curTime) && !n.CheckWasSentToday(curTime) {
			logger.Get("main").Debugf("sending notification for %+v", n)
			sendNotificationForUser(n, bot)
			n.LastTimeActivated = curTime.Unix()
			daily_notifications.Save(n)
		}
	}
}

func sendNotificationForUser(n daily_notifications.DailyNotificationConfig, bot *tgbotapi.BotAPI) {
	tasks, err := wservice.GetUserTodayTasks(n.UserID)
	if err != nil {
		logger.Get("main").Critical(err)
		return
	}

	if len(tasks) == 0 {
		logger.Get("main").Infof("user %v has 0 tasks for today, aborting notification send", n.UserID)
		return
	}

	logger.Get("main").Infof("sending notification user %v", n.UserID)
	msgText := "Hi! Here are your tasks for today:\n"
	for _, t := range tasks {
		msgText += "▫ " + t.Title + "\n"
	}

	msg := tgbotapi.NewMessage(n.ChatID, msgText)
	bot.Send(msg)

}
