package main

import (
	"github.com/JahidNishat/BusTicketChecker/busbdChckr/notifier"
	"github.com/JahidNishat/BusTicketChecker/busbdChckr/stations"
	"github.com/spf13/viper"
	"log"
)

func init() {
	stations.ProcessStationMap()
}

func main() {
	viper.SetConfigFile("./config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	//go job.StartTicker()
	log.Println("Server running...")
	notifier.ServeTgBot()
}
