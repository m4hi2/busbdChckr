package utils

import (
	"strings"
)

// LevenshteinDistance compares two strings and returns the levenshtein distance between them.
func LevenshteinDistance(s, t string, ignoreCase bool) int {
	if ignoreCase {
		s = strings.ToLower(s)
		t = strings.ToLower(t)
	}
	d := make([][]int, len(s)+1)
	for i := range d {
		d[i] = make([]int, len(t)+1)
	}
	for i := range d {
		d[i][0] = i
	}
	for j := range d[0] {
		d[0][j] = j
	}
	for j := 1; j <= len(t); j++ {
		for i := 1; i <= len(s); i++ {
			if s[i-1] == t[j-1] {
				d[i][j] = d[i-1][j-1]
			} else {
				mn := d[i-1][j]
				if d[i][j-1] < mn {
					mn = d[i][j-1]
				}
				if d[i-1][j-1] < mn {
					mn = d[i-1][j-1]
				}
				d[i][j] = mn + 1
			}
		}

	}
	return d[len(s)][len(t)]
}

// FindClosestStation finds out the closest station from provided station list
func FindClosestStation(stationName string, stationList *[]string) string {
	minDistance := 10000
	var matched string

	for _, station := range *stationList {
		ld := min(LevenshteinDistance(station, stationName, true))
		suggestByLevenshtein := ld <= 3
		suggestByPrefix := strings.HasPrefix(strings.ToLower(station), strings.ToLower(stationName))
		if suggestByLevenshtein && ld < minDistance || suggestByPrefix {
			minDistance = ld
			matched = station
		}
	}
	return matched
}

// GetClosestStation returns the close matched station if found, returns the same name otherwise
func GetClosestStation(stationName string, stationList *[]string) string {
	matched := FindClosestStation(stationName, stationList)
	if matched != "" {
		return matched
	}
	return stationName
}
