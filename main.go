package main

import (
	"github.com/fahimimam/busbdChckr/tickerv1"
	"github.com/fahimimam/busbdChckr/tickerv2"
)

const (
	runner = "v2"
)

func main() {
	if runner == "v2" {
		tickerv2.RunTickerV2()
	} else {
		tickerv1.RunTickerV1()
	}
}
