package job

import (
	"github.com/m4hi2/busbdChckr/proccessor"
	"github.com/spf13/viper"
	"log"
	"time"
)

// TODO: Ticker will be added  later

func StartTicker() {
	tickerDuration := viper.GetDuration("ticker.duration")
	if tickerDuration <= 0 {
		tickerDuration = 1 * time.Hour
	}
	ticker := time.NewTicker(tickerDuration)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			log.Println("######################### Started Fetching #########################")

			proccessor.GetUserData()

			log.Printf("######################### Finished Fetching #########################\n\n\n")
		}
	}
}
