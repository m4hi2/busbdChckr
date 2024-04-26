package main

import (
	"github.com/m4hi2/busbdChckr/db"
	"github.com/m4hi2/busbdChckr/job"
	"github.com/m4hi2/busbdChckr/notifier"
	"github.com/m4hi2/busbdChckr/stations"
	"github.com/spf13/viper"
	"log"
)

func InitAll() {
	stations.ProcessStationMap()

	viper.SetConfigFile("./config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	db.ConnectDB()
}

func main() {
	InitAll()

	go job.StartTicker()
	log.Println("Server running...")
	notifier.ServeTgBot()
}
