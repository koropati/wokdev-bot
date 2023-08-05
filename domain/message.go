package domain

import (
	"fmt"
	"strconv"
	"wokdev-bot/helper"
	"wokdev-bot/infra"
)

type MessageService interface {
	SendMessage(userID int64, text string)
	SendMaintenanceMessage(userID int64)
	SendMaintenanceMessageWebsite(userID int64, applicationName string, urlApp string)
	SendRunningMessage(userID int64)
	SendRunningMessageWebsite(userID int64, applicationName string, urlApp string)
	SendHelpMessage(userID int64)
	SendInfoID(userID int64, fromID int)
}

type messageService struct {
	telegram infra.TelegramService
}

func NewMessageService(telegram infra.TelegramService) MessageService {
	return &messageService{
		telegram: telegram,
	}
}

func (svc *messageService) SendMaintenanceMessage(userID int64) {
	// Ganti pesan sesuai kebutuhan
	text := "Website sedang dalam maintenance. Mohon maaf atas ketidaknyamanan ini."
	svc.telegram.SendMessage(userID, text)
}

func (svc *messageService) SendMaintenanceMessageWebsite(userID int64, applicationName string, urlApp string) {
	// Ganti pesan sesuai kebutuhan
	text := "Backend situs " + applicationName + " sedang dalam maintenance. Situs " + urlApp + " mungkin tidak dapat berjalan dengan baik, Mohon maaf atas ketidaknyamanan ini."
	svc.telegram.SendMessage(userID, text)
}

func (svc *messageService) SendRunningMessage(userID int64) {
	// Ganti pesan sesuai kebutuhan
	text := "Website sedang berjalan dengan baik. Terimakasih!"
	svc.telegram.SendMessage(userID, text)
}

func (svc *messageService) SendRunningMessageWebsite(userID int64, applicationName string, urlApp string) {
	// Ganti pesan sesuai kebutuhan
	text := "Backend situs " + applicationName + " sedang berjalan dengan baik. Anda bisa mengaksesnya melalui link : " + urlApp + " Terimakasih!"
	svc.telegram.SendMessage(userID, text)
}

func (svc *messageService) SendMessage(userID int64, text string) {
	svc.telegram.SendMessage(userID, text)
}

func (svc *messageService) SendHelpMessage(userID int64) {
	pesanBantuan := fmt.Sprintf(`%s %s, ada yang bisa saya bantu?, berikut merupakan list perintah yang bisa anda gunakan: `, helper.GetRandomGreetingOpening(), helper.GetGreeting())
	commandHelp := "/help\t\t: melihat bantuan"
	commandStatus := "/status_{kode_aplikasi}\t\t: melihat status aplikasi"
	commandMyID := "/my_id\t\t: melihat ID Telegram Anda"
	pesanPenutup := "Terimakasih, semoga membantu."

	message := pesanBantuan + "\n\n" + commandHelp + "\n" + commandStatus + "\n" + commandMyID + "\n\n" + pesanPenutup

	svc.telegram.SendMessage(userID, message)

}

func (svc *messageService) SendInfoID(userID int64, fromID int) {
	text := "ID Telegram anda adalah : " + strconv.Itoa(fromID)
	svc.telegram.SendMessage(userID, text)
}
