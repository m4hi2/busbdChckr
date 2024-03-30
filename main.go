package main

import (
	"fmt"
	"github.com/JahidNishat/BusTicketChecker/busbdChckr/job"
	"github.com/JahidNishat/BusTicketChecker/busbdChckr/notifier"
	"github.com/JahidNishat/BusTicketChecker/busbdChckr/stations"
)

func init() {
	stations.ProcessStationMap()
}

func main() {
	go job.StartTicker()
	fmt.Println("Server running...")
	notifier.ServeTgBot()

}
