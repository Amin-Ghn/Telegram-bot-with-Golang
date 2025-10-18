package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gorm.io/gorm"
	"log"
)

func Route(bot *tgbotapi.BotAPI, update tgbotapi.Update, db *gorm.DB) {

	if update.Message == nil {
		return
	}
	var user User
	db.Where("chat_id = ?", update.Message.Chat.ID).First(&user)
	if user.ID == 0 {

		AskLanguage(bot, update.Message.Chat.ID)
		return
	}
	if user.Language == "" {
		handleLanguageSelection(bot, update, db)
		return
	}
	HandleMainMenu(bot, update, db, user)
}
