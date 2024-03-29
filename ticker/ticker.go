package ticker

import (
	"errors"
	"github.com/fahimimam/busbdChckr/routeInformation"
	avbusinfo "github.com/fahimimam/busbdChckr/routeInformation/busInformation"
	"github.com/fahimimam/busbdChckr/routeInformation/models"
	"github.com/fahimimam/busbdChckr/stations"
	"log"
	"strings"
	"time"
)

const StructureTypeBus = "BUS"

func RetrieveInformation(source, destination, date string) {
	sourceID := stations.StationCodeToStationID[strings.ToLower(source)]
	destID := stations.StationCodeToStationID[strings.ToLower(destination)]
	// How am I getting data
	data := avbusinfo.RequestPld{
		Date:          date,
		FromStationId: sourceID,
		ToStationId:   destID,
	}
	// Initial fetch
	prevData, err := fetchBusInfo(data, StructureTypeBus)
	if err != nil {
		log.Println("Error fetching initial data:", err)
	}
	log.Println("Fetched Data from server...")
	// Start a goroutine to periodically check for updates
	ticker := time.NewTicker(1 * time.Hour) // Adjust interval as needed
	defer ticker.Stop()
	log.Println("Ticker Running...")
	for {
		select {
		case <-ticker.C:
			updatedData, err := fetchBusInfo(data, StructureTypeBus)
			if err != nil {
				log.Println("Error fetching updated data:", err)
				continue
			}

			if hasUpdates(prevData, updatedData) {
				log.Printf("Data has been updated - \n%+v", updatedData)
			} else {
				log.Println("Data has not been updated")
			}

			prevData = updatedData
		}
	}
}

// Function to fetch bus information
func fetchBusInfo(data avbusinfo.RequestPld, structureType string) ([]*models.NotificationPld, error) {
	switch structureType {
	case StructureTypeBus:
		return routeInformation.GetBusInfo(data)
	default:
		return nil, errors.New("enter valid travel medium")
	}

}

// Function to compare if there are updates in data
func hasUpdates(prev, updated []*models.NotificationPld) bool {
	return !isEqual(prev, updated)
}

// Function to compare if two response payloads are equal
func isEqual(a, b []*models.NotificationPld) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i].RouteId != b[i].RouteId ||
			a[i].CoachNo != b[i].CoachNo ||
			a[i].RouteName != b[i].RouteName ||
			a[i].AvailableSeats != b[i].AvailableSeats ||
			a[i].StartCounter != b[i].StartCounter ||
			!equalStringPointer(a[i].ArrivaleTime, b[i].ArrivaleTime) ||
			a[i].DepartureTime != b[i].DepartureTime ||
			a[i].CompanyName != b[i].CompanyName ||
			a[i].Fare != b[i].Fare {
			return false
		}
	}
	return true
}

// Function to compare two string pointers
func equalStringPointer(a, b *string) bool {
	if a == nil || b == nil {
		return a == b
	}
	return *a == *b
}
