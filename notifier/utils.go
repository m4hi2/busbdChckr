package notifier

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/m4hi2/busbdChckr/db"
	"github.com/m4hi2/busbdChckr/db/models"
	"github.com/m4hi2/busbdChckr/db/repos"
	rm "github.com/m4hi2/busbdChckr/routeInformation/models"
	"gorm.io/gorm"
	"log"
	"reflect"
)

func StringifyStruct(data interface{}) string {
	val := reflect.ValueOf(data)
	typ := reflect.TypeOf(data)

	if val.Kind() != reflect.Struct {
		return "Not a struct"
	}

	var str string
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		str += fmt.Sprintf("%s: %v\n", fieldType.Name, field.Interface())
	}

	return str
}

func HashData(str string) string {
	sha1Hash := sha1.New()
	sha1Hash.Write([]byte(str))
	hashBytes := sha1Hash.Sum(nil)
	sha1String := hex.EncodeToString(hashBytes)
	return sha1String
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

func ProcessData(busInfos []*rm.ResponsePld, seat int, chatID int64) error {
	u := repos.UserStore{db.ConnectDB()}

	message := fmt.Sprintf("Total Available Coach: %v\n Total Available Seats: %v\n", len(busInfos), seat)
	bot, err := GetTelegramBot()
	if err != nil {
		log.Println("Fetching Telegram Bot Error: ", err)
		return err
	}

	bot.SendMessage(chatID, message)

	for _, busInfo := range busInfos {
		stringifyData := StringifyStruct(*busInfo)
		log.Println(stringifyData)

		hashData := HashData(stringifyData)
		log.Println(hashData)

		logData, err := u.FindHash(hashData)
		if errors.Is(err, gorm.ErrRecordNotFound) && logData == nil {
			err = PublishToUsers(chatID, stringifyData)
			if err != nil {
				log.Println("Failed while sending the update")
				continue
			}

			// Insert new hash data
			err = u.InsertData(&models.Log{
				ChatID: chatID,
				Hash:   hashData,
				Data:   stringifyData,
			})
			if err != nil {
				log.Println("Failed while inserting data into database: ", err)
				return err
			}
		}
	}
	return nil
}
