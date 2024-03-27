package job

import (
	"github.com/spf13/viper"
	"log"
	"time"
)

func StartTicker() {
	ticker := time.NewTicker(viper.GetDuration("ticker.duration"))

	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			log.Println("######################### Started Fetching #########################")

			log.Println("######################### Finished Fetching #########################\n\n")
		}
	}
}
