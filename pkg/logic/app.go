package logic

import (
	botapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func BotWork(bot *botapi.BotAPI){
	bot.Debug = true

	updCfg := botapi.NewUpdate(0)
	// Таймер проверки обновлений
	updCfg.Timeout = 30
	updates := bot.GetUpdatesChan(updCfg)

	for update := range updates{
		if update.Message == nil{
			continue
		}

		// Новое сообщение бота
		msg := botapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		// Говорим, что сообщение от бота является ответом, значит создает ответ, а не просто выводит текст
		msg.ReplyToMessageID = update.Message.MessageID

		if _,err:=bot.Send(msg);err!=nil{
			panic(err)
		}
	}
}