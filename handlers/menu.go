package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"gorm.io/gorm"
)

func ShowMainMenu(bot *tgbotapi.BotAPI, chatID int6, lang string ) {
	var text string
	switch lang {
	case "fa":

		text = "منو اصلی:\nمحصولات\nدرباره ما\nنمایندگان\nتماس با ما"
	case "ar":
		text = "القائمه الرییسیه\nالمنتجات\nمعلومات عنا\nنمایندگان\nاتصل بنا"
	default:
		text = "Main Menu:\nProducts\nAbout Us:\nRepresentatives\nContact Us:"

	}
msg := tgbotapi.NewMessage(chatID, text)
msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(getButtonText("products", lang)),
		tgbotapi.NewKeyboardButton(getButtonText("About Us", lang)),
		tgbotapi.NewKeyboardButton(getButtonText("Contact Us", lang)),
		tgbotapi.NewKeyboardButton(getButtonText("Representatives", lang)),

		),

)
bot.Send(msg)

}
func getButtonText(key string, lang string) string {
	translations := map[string]map[string]string{
		"products" :{
			"fa": "محصولات",
			"ar": "المنتجات",
			"en": "products",
		},
		"AboutUs": {
			"fa" : "درباره ما",
			"ar" : "معلومات عنا",
			"en" : "About Us",
		},
		"ContactUs": {
			"fa" : "ارتباط با ما",
			"ar" : "اتصل بنا",
			"en" : "Contact Us",
		},
		"Representatives" : {
			"fa" : "نمایندگان",
			"ar" : "الممثلین",
			"en" : "Representatives",
		}
		if trans, ok := translations[key]; ok {
			if text, ok := trans[lang]; ok {
				return text
			}

		}
	}
}