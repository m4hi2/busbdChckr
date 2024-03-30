package job

import (
	"github.com/spf13/viper"
	"log"
	"time"
)

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

			log.Printf("######################### Finished Fetching #########################\n\n\n")
		}
	}
}
