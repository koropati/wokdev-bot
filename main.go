package main

import (
	"log"

	"wokdev-bot/app"
	"wokdev-bot/domain"
	"wokdev-bot/infra"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	token := "xxxxxccccccccxxxxx"

	messageSvc := domain.NewMessageService(infra.NewTelegramService(initBot(token)))
	bot, err := app.NewBot(token, messageSvc)
	if err != nil {
		log.Panic(err)
	}
	log.Println("Aplikasi Berjalan")
	bot.Run()
}

func initBot(token string) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	return bot
}
