package server

import (
	"log"

	botapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"bot/pkg/logic"
)

func Launch(api string) error{
	bot, err := botapi.NewBotAPI(api)
	if err!=nil{
		log.Fatalf("%s", err.Error())
	}

	logic.BotWork(bot)

	return err
}