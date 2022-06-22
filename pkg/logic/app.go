package logic

import (
	"log"

	botapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)


var infoboard = botapi.NewInlineKeyboardMarkup(
	botapi.NewInlineKeyboardRow(
		botapi.NewInlineKeyboardButtonURL("Github", "https://github.com/Peygy"),
	),
	botapi.NewInlineKeyboardRow(
		botapi.NewInlineKeyboardButtonData("WebTutorial", "Скоро будет"),
	),
	botapi.NewInlineKeyboardRow(
		botapi.NewInlineKeyboardButtonData("WebOrginizer", "Скоро будет"),
	),
)


func Receiver(bot *botapi.BotAPI){
	bot.Debug = true

	updCfg := botapi.NewUpdate(0)
	updCfg.Timeout = 60

	updates := bot.GetUpdatesChan(updCfg)
	for update := range updates{
		if update.Message != nil{
			if !update.Message.IsCommand(){
				MessageUpd(bot, update)
			}else if update.Message.IsCommand(){
				CommandUpd(bot, update)
			}
		}else if update.CallbackQuery != nil{
			CallbackUpd(bot, update)
		}
	}
}


func MessageUpd(bot *botapi.BotAPI, update botapi.Update){
	msg := botapi.NewMessage(update.Message.Chat.ID, "Это что ТЕКСТ! Я понимаю только команды!")

	if _,err := bot.Send(msg); err != nil{
		log.Fatalf("%s", err.Error())
		msg.Text = "Технические проблемы, попробуйте снова;)"
		bot.Send(msg)
	}
}


func CommandUpd(bot *botapi.BotAPI, update botapi.Update){
	msg := botapi.NewMessage(update.Message.Chat.ID, "")
	userName := update.SentFrom().UserName

	switch update.Message.Command(){
	case "start":
		msg.Text = "Привет "+userName+" !\nПока что наши сайты разрабатываются, "+
		"но скоро Вы сможете с ними ознакомиться!\n"+
		"Воспользуйся командой /info, чтобы получить нужную информацию(Санечка сниаешь?)"

	case "info":
		msg.Text = "Наши проекты и ссылки к ним"
		msg.ReplyMarkup = infoboard

	default:
		msg.Text = "Охх... я не знаю такой команды (° ͜ʖ°)"
	}

	if _,err := bot.Send(msg); err != nil{
		log.Fatalf("%s", err.Error())
		msg.Text = "Технические проблемы, попробуйте снова;)"
		bot.Send(msg)
	}
}


func CallbackUpd(bot *botapi.BotAPI, update botapi.Update){
	callback := botapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
	msg := botapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "")
	if _,err := bot.Request(callback); err != nil{
		log.Fatalf("%s", err.Error())
		msg.Text = "Технические проблемы, попробуйте снова;)"
		bot.Send(msg)
	}

	msg.Text = update.CallbackQuery.Data
	if _,err := bot.Send(msg); err != nil{
		log.Fatalf("%s", err.Error())
		msg.Text = "Технические проблемы, попробуйте снова;)"
		bot.Send(msg)
	}
}