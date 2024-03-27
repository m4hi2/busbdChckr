package notifier

import (
	"fmt"
	"strings"
	"time"
)

func (bot *TelegramBot) StartCMD(chatID int64) {
	bot.SendMessage(chatID, "Welcome to Bus Ticket Availability Checker! Please use "+
		"/check <source> <destination> <date> command to check ticket availability.")
}

func (bot *TelegramBot) CheckCMD(messageText string, chatID int64) {
	// Parse user input to get source, destination, and date
	parts := strings.Split(messageText, " ")
	if len(parts) != 4 {
		bot.SendMessage(chatID, "Invalid command usage. Please use /check <source> <destination> <date>")
		return
	}
	source := parts[1]
	destination := parts[2]
	dateStr := parts[3]
	_, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		bot.SendMessage(chatID, "Invalid date format. Please use YYYY-MM-DD")
		return
	}

	// TODO Have to call Bus Ticket Api
	// Call for Bus Ticket Api
	//data, err := Processor(source, destination, dateStr)
	if err != nil {
		bot.SendMessage(chatID, "Cannot Fetch Data")
		return
	}

	message := fmt.Sprintf("Tickets from %s to %s on %s are available: %t", source, destination, dateStr)
	bot.SendMessage(chatID, message)
	return
}

func (bot *TelegramBot) HelpCMD(chatID int64) {
	bot.SendMessage(chatID, "Available commands:\n"+
		"/check Source Destination Date (YYYY-MM-DD) - Check bus ticket availability\n"+
		"/help - Show this help message")
}
