package job

import "github.com/m4hi2/busbdChckr/routeInformation/models"

const StructureTypeBus = "BUS"

// TODO : Rethink or discard in future depending on the needs.
func hasUpdates(prev, updated []*models.ResponsePld) bool {
	return !isEqual(prev, updated)
}

// Function to compare if two response payloads are equal
func isEqual(a, b []*models.ResponsePld) bool {
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
