package proccessor

import (
	"github.com/m4hi2/busbdChckr/db"
	"github.com/m4hi2/busbdChckr/db/repos"
	"github.com/m4hi2/busbdChckr/notifier"
	"github.com/m4hi2/busbdChckr/routeInformation"
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
		busInfos, seat, err := routeInformation.GetBusInfo(user.Source, user.Destination, user.Date)
		if err != nil {
			log.Println("Fetching Bus Info Error: ", err)
			return err
		}

		err2 := notifier.ProcessData(busInfos, seat, user.ChatID)
		if err2 != nil {
			return err2
		}
	}

	return nil
}
