package handlers

import (
	tgbotapi " github.com/go-telegram-bot-api/telegram-bot-api"
	"gorm.io/gorm"
)

func AskLanguage(bot *tgbotapi.BotAPI, chatID int64) {

	msg := tgbotapi.NewMessage(chatID, "Choose your language:\nزبان خود را انتخاب کنید: \n اختر لغتک:")
	msg.replyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("English"),
			tgbotapi.NewKeyboardButton("فارسی"),
			tgbotapi.NewKeyboardButton("العربیه"),
		),
	)
	bot.Send(msg)
}
