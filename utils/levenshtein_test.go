package utils

import (
	"fmt"
	"github.com/m4hi2/busbdChckr/stations"
	"testing"
)

func Test_LevenshteinDistance(t *testing.T) {
	input := "chapainobabganj"

	stations.ProcessStationMap()
	matched := FindClosestStation(input, &stations.StationNames)

	if matched != "" {
		fmt.Println("Found a match : ", matched)
	}
}
