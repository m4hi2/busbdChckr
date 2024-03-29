package main

import (
	"github.com/fahimimam/busbdChckr/stations"
	"github.com/fahimimam/busbdChckr/ticker"
)

func init() {
	stations.ProcessStationMap()
}

func main() {
	ticker.RetrieveInformation("dhaka", "coxs-bazar", "2024-04-04")
}
