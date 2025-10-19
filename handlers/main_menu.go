package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gorm.io/gorm"
)

func HandleMainMenu(bot *tgbotapi.BotAPI, update tgbotapi.Update, db *gorm.db, user User) {
	text := update.Message.Text
	switch text {
	case getButtonText("Products", user.Language):
		HandleProducts(bot, update.Message.Chat.ID, user.Language, db)
	case getButtonText("About Us", user.Language):
		HandleAboutUs(bot, update.Message.chat.ID, user.Language)
	case getButtonText("Contact Us", user.Language):
		HandleContactUs(bot, update.Message.Chat.ID, user.Language)
	case getButtonText("Representatives", user.Language):
		HandleRepresentatives(bot, update.Message.Chat.ID)
	}
}
