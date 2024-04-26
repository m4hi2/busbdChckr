package stations

import (
	"encoding/json"
	"fmt"
	"os"
)

type Station struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type CodeToID map[string]string

var StationCodeToStationID CodeToID
var StationNames []string

func ProcessStationMap() {
	file, err := os.Open("stations/locations.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Decode the JSON file into a slice of Station structs
	var stations []Station
	err = json.NewDecoder(file).Decode(&stations)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	StationCodeToStationID = make(CodeToID)
	// Create a hashmap to store the mapping from code to ID
	for _, station := range stations {
		StationCodeToStationID[station.Code] = station.ID
		StationNames = append(StationNames, station.Code)
	}
}
