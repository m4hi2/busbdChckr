package notifier

import (
	"errors"
	"fmt"
	"github.com/m4hi2/busbdChckr/db"
	"github.com/m4hi2/busbdChckr/db/models"
	"github.com/m4hi2/busbdChckr/db/repos"
	rm "github.com/m4hi2/busbdChckr/routeInformation/models"
	"gorm.io/gorm"
	"log"
)

func ProcessData(busInfos []*rm.ResponsePld, chatID int64, source, destination, date string) error {
	u := repos.UserStore{db.ConnectDB()}
	var availableCoach, availableSeat int

	for _, busInfo := range busInfos {
		hashData := HashData(busInfo)
		log.Println("Hashed Data: ", hashData)

		logData, err := u.FindHash(hashData)
		if errors.Is(err, gorm.ErrRecordNotFound) && logData == nil {
			responseData := ResponseStruct(busInfo)

			err = PublishToUsers(chatID, responseData)
			if err != nil {
				log.Println("Failed while sending the update")
				continue
			}

			// Insert new hash data
			err = u.InsertData(&models.Log{
				ChatID: chatID,
				Hash:   hashData,
				Data:   responseData,
			})
			if err != nil {
				log.Println("Failed while inserting data into database: ", err)
				return err
			}

			if busInfo.AvailableSeats > 0 {
				availableSeat += busInfo.AvailableSeats
				availableCoach++
			}
		}
	}

	message := fmt.Sprintf("*From:* _%v_\n*To:* _%v_\n*Date:* _%v_\n"+
		"*Total New Available Coach:* _%v_\n*Total New Available Seats:* _%v_\n",
		source, destination, date, availableCoach, availableSeat)
	bot, err := GetTelegramBot()
	if err != nil {
		log.Println("Fetching Telegram Bot Error: ", err)
		return err
	}

	if availableSeat > 0 && availableCoach > 0 {
		bot.SendMessage(chatID, message)
	}

	return nil
}

func PublishToUsers(chatID int64, stringifyData string) error {
	bot, err := GetTelegramBot()
	if err != nil {
		log.Println("Error getting Telegram bot:", err)
		return err
	}

	bot.SendMessage(chatID, stringifyData)

	return nil
}
