package proccessor

import (
	"github.com/m4hi2/busbdChckr/db"
	"github.com/m4hi2/busbdChckr/db/repos"
	"github.com/m4hi2/busbdChckr/notifier"
	"github.com/m4hi2/busbdChckr/routeInformation"
	"github.com/m4hi2/busbdChckr/utils"
	"log"
)

func GetUserData() error {
	u := repos.UserStore{db.ConnectDB()}

	users, err := u.GetAll()
	if err != nil {
		log.Println("Error getting all data", err)
		return err
	}
	if len(users) == 0 {
		log.Println("No users found")
		return nil
	}

	for _, user := range users {
		if utils.IsDateExpired(user.Date) {
			log.Println("Date expired")

			err := u.DeleteUser(user)
			if err != nil {
				log.Println("Error deleting user", err)
			}

			continue
		}

		busInfos, _, err := routeInformation.GetBusInfo(user.Source, user.Destination, user.Date)
		if err != nil {
			log.Println("Fetching Bus Info Error: ", err)
			return err
		}

		err2 := notifier.ProcessData(busInfos, user.ChatID)
		if err2 != nil {
			return err2
		}
	}

	return nil
}
