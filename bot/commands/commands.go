package commands

import (
	"errors"
	"fmt"

	"gopkg.in/telegram-bot-api.v4"
)

var availableCommands = map[string]func(*tgbotapi.Message, *tgbotapi.BotAPI){
	"start": start,
}

//GetHandler returns command handler by command name
func GetHandler(command string) (func(*tgbotapi.Message, *tgbotapi.BotAPI), error) {
	cmd := availableCommands[command]
	if cmd == nil {
		errorMessage := fmt.Sprintf("unknown command %v", command)
		return nil, errors.New(errorMessage)
	}

	return cmd, nil
}
