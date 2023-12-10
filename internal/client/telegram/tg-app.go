package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"my-app/internal/config"
)

type TgApi struct {
	bot    *tgbotapi.BotAPI
	config config.Config
}

var tgApi TgApi

func init() {
	cfg, err := config.FromEnv()
	if err != nil {
		log.Fatalf("Failed to read configuration: %v", err)
		return
	}

	bot, err := tgbotapi.NewBotAPI(cfg.TelegramBotConfig.ApiToken)

	if err != nil {
		log.Fatalf("Failed to create Telegram bot: %v", err)
		return
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	tgApi = TgApi{bot: bot, config: cfg}
}

func Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := tgApi.bot.GetUpdatesChan(u)

	for update := range updates {
		handleUpdate(update)
	}
}

func handleUpdate(update tgbotapi.Update) {

	message := update.Message

	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, "new message")

	msg.ParseMode = tgbotapi.ModeMarkdown
	msg.ReplyToMessageID = message.MessageID

	_, err := tgApi.bot.Send(msg)

	if err != nil {
		return
	}
}
