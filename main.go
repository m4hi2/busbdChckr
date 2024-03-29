package main

import (
	"github.com/fahimimam/busbdChckr/stations"
	"github.com/fahimimam/busbdChckr/ticker"
)

const (
	runner = "v2"
)

func init() {
	stations.ProcessStationMap()
}

func main() {
	ticker.RetrieveInformation("dhaka", "coxs-bazar", "2024-04-04")
}
