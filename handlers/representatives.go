package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gorm.io/gorm"
)

func HandleRepresentatives(bot *tgbotapi.BotAPI, chatID int64, lang string, db *gorm.DB) {
	var text string

	switch lang {
	case "fa":
		text = "نمایندگان: \n\n" +
			"تهران: خیابان\n" +
			"091211111"
	case "ar":
		text = "ممثلونا: \n\n" +
			"تهران: شارع\n" +
			"091211111"
	default:
		text = "Representatives: \n\n" +
			"Tehran: st.\n" +
			"091211111"

	}
	msg := tgbotapi.NewMessage(chatID, text)
	bot.Send(msg)

	ShowMainMenu(bot, chatID, lang)

}
