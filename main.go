package main

import (
	"github.com/m4hi2/busbdChckr/db"
	"github.com/m4hi2/busbdChckr/stations"
	"github.com/spf13/viper"
	"log"
)

func init() {
	stations.ProcessStationMap()

	viper.SetConfigFile("./config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	err := db.DoPersistConnect()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
}

func main() {
	//go job.StartTicker()
	log.Println("Server running...")
	notifier.ServeTgBot()
}
