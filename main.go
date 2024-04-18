package main

import (
	"github.com/m4hi2/busbdChckr/notifier"
	"github.com/m4hi2/busbdChckr/stations"
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
