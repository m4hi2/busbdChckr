package notifier

import (
	"github.com/m4hi2/busbdChckr/db"
	"github.com/m4hi2/busbdChckr/db/models"
	"github.com/m4hi2/busbdChckr/db/repos"
	"github.com/m4hi2/busbdChckr/routeInformation"
	"github.com/m4hi2/busbdChckr/stations"
	"github.com/m4hi2/busbdChckr/utils"
	"strings"
)

func (bot *TelegramBot) StartCMD(chatID int64) {
	bot.SendMessage(chatID, "Welcome to Bus Ticket Availability Checker! Please use "+
		"/check <source> <destination> <date>[YYYY-MM-DD] command to check ticket availability.")
}

func (bot *TelegramBot) CheckCMD(messageText string, chatID int64) {
	// Parse user input to get source, destination, and date
	parts := strings.Split(messageText, " ")
	if len(parts) != 4 {
		bot.SendMessage(chatID, "Invalid command usage. Please use /check <source> <destination> <date>[YYYY-MM-DD] command to check ticket availability.")
		return
	}
	source := parts[1]
	destination := parts[2]
	dateStr := parts[3]

	//expired := utils.IsDateExpired(dateStr)

	if utils.IsDateExpired(dateStr) {
		bot.SendMessage(chatID, "Invalid date or date format. Please use YYYY-MM-DD or ensure it's not a past date.")
		return
	}

	source = utils.GetClosestStation(source, &stations.StationNames)
	destination = utils.GetClosestStation(destination, &stations.StationNames)

	resPld, _, err := routeInformation.GetBusInfo(source, destination, dateStr)
	if err != nil || resPld == nil {
		bot.SendMessage(chatID, "Invalid Source or Destination. Please use valid Source or Destination.")
		return
	}

	err = ProcessData(resPld, chatID)
	if err != nil {
		bot.SendMessage(chatID, "Cannot Fetch Data")
		return
	}

	u := repos.UserStore{DB: db.ConnectDB()}
	_ = u.CreateUser(&models.User{
		ChatID:      chatID,
		Source:      source,
		Destination: destination,
		Date:        dateStr,
	})

	return
}

func (bot *TelegramBot) HelpCMD(chatID int64) {
	bot.SendMessage(chatID, "Available commands:\n"+
		"/check Source Destination Date (YYYY-MM-DD) - Check bus ticket availability\n"+
		"/help - Show this help message")
}
