package main

import (
	"Coreet-Telegram-Bot/database"
	"Coreet-Telegram-Bot/handlers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}
	db := database.InitDB()
	defer db.Close()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		handlers.Route(bot, update, db)

	}
}
