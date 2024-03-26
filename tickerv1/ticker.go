package tickerv1

import (
	"log"
	"time"
)

func RunTickerV1() {
	t := time.NewTicker(2 * time.Second)

	for _ = range t.C {

		x := GetBusInfo()

		for _, seat := range x {
			log.Println(seat)
			log.Println("......")
		}

	}
}
