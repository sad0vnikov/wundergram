package main

import (
	"os"

	"github.com/sad0vnikov/wundergram/bot"
	"github.com/sad0vnikov/wundergram/bot/commands"
	"github.com/sad0vnikov/wundergram/callbacklistener"
)

func main() {
	token := os.Getenv("TELEGRAM_TOKEN")
	if len(token) == 0 {
		panic("telegaram token is not set")
	}
	go callbacklistener.Init()

	dialogTree := commands.BuildConversationTree()
	bot.Init(token, dialogTree)
}