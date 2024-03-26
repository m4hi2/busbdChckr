package tickerv2

import (
	"fmt"
	"github.com/m4hi2/busbdChckr/businfo"
	"github.com/m4hi2/busbdChckr/businfo/models"
	"log"
	"strings"
	"time"
)

const StructureTypeBus = "BUS"

func RunTickerV2() {
	source := "Dhaka"
	source = strings.ToLower(source)
	destination := "khulna"
	destination = strings.ToLower(destination)
	data := businfo.RequestPld{
		Date:          "2024-04-04",
		Identifier:    fmt.Sprintf("%s-to-%s", source, destination),
		StructureType: StructureTypeBus,
	}
	// Initial fetch
	prevData, err := fetchBusInfo(data)
	if err != nil {
		log.Println("Error fetching initial data:", err)
	}

	// Start a goroutine to periodically check for updates
	ticker := time.NewTicker(1 * time.Hour) // Adjust interval as needed
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			updatedData, err := fetchBusInfo(data)
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
func fetchBusInfo(data businfo.RequestPld) ([]*models.NotificationPld, error) {
	return businfo.GetBusInfoV2(data)
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
