package handlers

import (
	"Go-CoreetTelegramBot/database"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gorm.io/gorm"
	"strconv"
)

func Handleproducts(bot *tgbotapi.BotAPI, chatID int64, lang string, db *gorm.DB) {
	var products []Product
	db.Limit(10).Find(&products)

	if len(products) == 0 {
		msg := tgbotapi.NewMessage(chatID, "No products found")
		bot.Send(msg)
		return
	}
	var text string
	switch lang {
	case "fa":

		text = "\n\nمحصولات ما"
	case "ar":
		text = "منتجاتنا\n\n"
	default:
		text = "Our Products:\n\n"
	}
	for i, product := range products {
		switch lang {
		case "fa":
			text += strconv.Itoa(i+1) + ". " + product.NameFA + "\n"
		case "ar":
			text += strconv.Itoa(i+1) + ". " + product.NameAR + "\n"
		default:
			text += strconv.Itoa(i+1) + ". " + product.NameEN + "\n"
		}
		msg := tgbotapi.NewMessage(chatID, text)

		keyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("بعدی", "products_page_2"),
			),
		)
		msg.ReplyMarkup = keyboard
		bot.Send(msg)
	}
}
