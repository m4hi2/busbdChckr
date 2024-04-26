package notifier

import (
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
	"log"
	"strings"
)

type TelegramBot struct {
	botAPI *tgbotapi.BotAPI
}

var tgBot *TelegramBot

func NewTelegramBot() (*TelegramBot, error) {
	token := viper.GetString("telegram.token")
	botAPI, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &TelegramBot{
		botAPI: botAPI,
	}, nil
}

func GetTelegramBot() (*TelegramBot, error) {
	var err error
	if tgBot == nil {
		tgBot, err = NewTelegramBot()
	}

	return tgBot, err
}

func (bot *TelegramBot) HandleIncomingMessage(update *tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	messageText := update.Message.Text

	log.Printf("Chat ID: %v -> Msg: %v\n", chatID, messageText)

	switch {
	case strings.HasPrefix(messageText, "/start"):
		bot.StartCMD(chatID)
	case strings.HasPrefix(messageText, "/check"):
		bot.CheckCMD(messageText, chatID)
	case strings.HasPrefix(messageText, "/help"):
		bot.HelpCMD(chatID)
	default:
		bot.SendMessage(chatID, "Invalid command. Please use /start or /check or /help.")
	}
}

func (bot *TelegramBot) SendMessage(chatID int64, message string) {
	msg := tgbotapi.NewMessage(chatID, message)
	msg.ParseMode = tgbotapi.ModeMarkdown
	_, err := bot.botAPI.Send(msg)
	if err != nil {
		log.Println("Error sending message:", err)
	}
}

func ServeTgBot() {
	telegramBot, err := GetTelegramBot()
	if err != nil {
		log.Println("Error initializing Telegram bot:", err)
		return
	}

	updateConfig := tgbotapi.NewUpdate(0)
	updates := telegramBot.botAPI.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message != nil {
			go telegramBot.HandleIncomingMessage(&update)
		}
	}
}
