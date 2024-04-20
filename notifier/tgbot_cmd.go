package notifier

import (
	"github.com/m4hi2/busbdChckr/db"
	"github.com/m4hi2/busbdChckr/db/models"
	"github.com/m4hi2/busbdChckr/db/repos"
	"github.com/m4hi2/busbdChckr/routeInformation"
	"github.com/m4hi2/busbdChckr/stations"
	"github.com/m4hi2/busbdChckr/utils"
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

	source = utils.GetClosestStation(source, &stations.StationNames)
	destination = utils.GetClosestStation(destination, &stations.StationNames)

	resPld, availableCount, err := routeInformation.GetBusInfo(source, destination, dateStr)
	if err != nil {
		bot.SendMessage(chatID, "Cannot Fetch Data")
		return
	}

	err = ProcessData(resPld, availableCount, chatID)
	if err != nil {
		bot.SendMessage(chatID, "Cannot Fetch Data")
		return
	}

	//for _, data := range resPld {
	//	message := StringifyStruct(*data)
	//	//fmt.Println(data)
	//	bot.SendMessage(chatID, message)
	//}

	//message := fmt.Sprintf("Total Available Coach: %v\n Total Available Seats: %v\n", len(resPld), availableCount)
	//bot.SendMessage(chatID, message)

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
