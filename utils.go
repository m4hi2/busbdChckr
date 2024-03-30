package main

import (
	"fmt"
	"time"
)

func GetBDTicketTimeFormat(t time.Time) string {
	formattedTime := t.Format("2006-01-02")
	fmt.Println(formattedTime)
	return formattedTime
}
