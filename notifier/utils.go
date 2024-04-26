package notifier

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
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

type ResponseData struct {
	CoachNo        string `json:"coachNo"`
	RouteName      string `json:"routeName"`
	AvailableSeats int    `json:"availableSeats"`
	StartCounter   string `json:"startCounter"`
	DepartureTime  string `json:"departureTime"`
	CompanyName    string `json:"companyName"`
	Fare           string `json:"fare"`
}

func ResponseStruct(busInfo *rm.ResponsePld) string {
	responseData := &ResponseData{
		CoachNo:        busInfo.CoachNo,
		RouteName:      busInfo.RouteName,
		AvailableSeats: busInfo.AvailableSeats,
		StartCounter:   busInfo.StartCounter,
		DepartureTime:  busInfo.DepartureTime,
		CompanyName:    busInfo.CompanyName,
		Fare:           busInfo.Fare,
	}

	return StringifyStruct(*responseData)
}

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

type HashedDataPld struct {
	RouteId     string `json:"routeId"`
	CoachNo     string `json:"coachNo"`
	RouteName   string `json:"routeName"`
	CompanyName string `json:"companyName"`
}

func HashData(busInfo *rm.ResponsePld) string {
	hashedDataPld := &HashedDataPld{
		RouteId:     busInfo.RouteId,
		CoachNo:     busInfo.CoachNo,
		RouteName:   busInfo.RouteName,
		CompanyName: busInfo.CompanyName,
	}

	jsonData, _ := json.Marshal(hashedDataPld)

	sha1Hash := sha1.New()
	sha1Hash.Write(jsonData)
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

type HashedData struct {
	RouteId       string `json:"routeId"`
	CoachNo       string `json:"coachNo"`
	RouteName     string `json:"routeName"`
	StartCounter  string `json:"startCounter"`
	DepartureTime string `json:"departureTime"`
	CompanyName   string `json:"companyName"`
	Fare          string `json:"fare"`
}

func ProcessData(busInfos []*rm.ResponsePld, chatID int64) error {
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

	message := fmt.Sprintf("Total New Available Coach: %v\n Total New Available Seats: %v\n", availableCoach, availableSeat)
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
