package main

import (
	"os"

	"github.com/sad0vnikov/wundergram/bot"
	"github.com/sad0vnikov/wundergram/bot/commands"
	"github.com/sad0vnikov/wundergram/callbacklistener"
	"github.com/sad0vnikov/wundergram/tasks"
)

func main() {
	token := os.Getenv("TELEGRAM_TOKEN")
	if len(token) == 0 {
		panic("telegaram token is not set")
	}
	go callbacklistener.Init()

	dialogTree := commands.BuildConversationTree()
	b := bot.Create(token)
	go b.Init(dialogTree)
	tasks.RunDailyNotifications(b)
}
