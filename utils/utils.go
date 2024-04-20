package utils

import (
	"log"
	"time"
)

func IsDateExpired(dateString string) bool {
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		log.Println("Error parsing date: ", err)
		return false
	}

	return date.Before(time.Now().Truncate(24 * time.Hour))
}
