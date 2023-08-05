package app

import (
	"log"
	"net/http"
	"strings"

	"wokdev-bot/domain"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	bot        *tgbotapi.BotAPI
	messageSvc domain.MessageService
}

func NewBot(token string, messageSvc domain.MessageService) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &Bot{
		bot:        bot,
		messageSvc: messageSvc,
	}, nil
}

func (b *Bot) Run() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		if update.Message != nil {
			PrintLogBot(update)
			if update.Message.IsCommand() && strings.Contains(update.Message.Command(), "status_") {
				b.sendStatus(update.Message.Chat.ID, update.Message.Command())
			} else {
				switch {
				case update.Message.IsCommand():
					switch update.Message.Command() {
					case "maintenance":
						b.messageSvc.SendMaintenanceMessage(update.Message.Chat.ID)
					case "running":
						b.messageSvc.SendRunningMessage(update.Message.Chat.ID)
					case "help":
						b.messageSvc.SendHelpMessage(update.Message.Chat.ID)
					case "my_id":
						b.messageSvc.SendInfoID(update.Message.Chat.ID, update.Message.From.ID)
					}
				}
			}

		}
	}
}

func PrintLogBot(update tgbotapi.Update) {
	username := update.Message.From.UserName
	firstName := update.Message.From.FirstName
	lastName := update.Message.From.LastName
	command := update.Message.Command()

	log.Printf("%s %s (%s) : Mengirim perintah: %s", firstName, lastName, username, command)
}

func GetURLApplication(applicationName string) (backendURL string, frontendURL string, isValid bool) {
	if applicationName == "senku" {
		return "https://attendance.wokdev.com", "https://senku-koropati.vercel.app/", true
	} else {
		return "", "", false
	}
}

func ExtractFullCommandStatus(fullCommand string) (applicationName string) {
	if strings.Contains(fullCommand, "status_") {
		applicationName = strings.ReplaceAll(fullCommand, "status_", "")
	} else {
		applicationName = fullCommand
	}
	return
}

func (b *Bot) sendStatus(chatID int64, fullCommand string) {
	// Ganti URL dengan URL situs yang ingin Anda cek statusnya
	applicationName := ExtractFullCommandStatus(fullCommand)
	url, frontEndURL, isValid := GetURLApplication(applicationName)
	if !isValid {
		b.messageSvc.SendMessage(chatID, "Maaf Kode Situs/Applikasi tidak terdaftar, silahkan menghubungi admin")
		return
	}
	resp, err := http.Get(url)
	if err != nil {
		b.messageSvc.SendMaintenanceMessageWebsite(chatID, applicationName, frontEndURL)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		b.messageSvc.SendRunningMessageWebsite(chatID, applicationName, frontEndURL)
	} else {
		b.messageSvc.SendMaintenanceMessageWebsite(chatID, applicationName, frontEndURL)
	}
}
