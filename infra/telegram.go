package infra

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

type TelegramService interface {
	SendMessage(userID int64, text string)
}

type telegramService struct {
	bot *tgbotapi.BotAPI
}

func NewTelegramService(bot *tgbotapi.BotAPI) TelegramService {
	return &telegramService{
		bot: bot,
	}
}

func (svc *telegramService) SendMessage(userID int64, text string) {
	msg := tgbotapi.NewMessage(userID, text)
	svc.bot.Send(msg)
}
