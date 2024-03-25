package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type BusInfo struct {
	Message interface{} `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    []struct {
		Id              string   `json:"id"`
		Date            string   `json:"date"`
		RouteId         string   `json:"routeId"`
		CoachConfigSKU  string   `json:"coachConfigSKU"`
		CoachConfigId   int      `json:"coachConfigId"`
		CoachNo         string   `json:"coachNo"`
		CoachType       string   `json:"coachType"`
		CoachSeatPlanId int      `json:"coachSeatPlanId"`
		SeatPlanSKU     *string  `json:"seatPlanSKU"`
		SeatTypesId     []string `json:"seatTypesId"`
		Fares           struct {
			EClass string `json:"E-Class,omitempty"`
			ECLASS string `json:"E-CLASS,omitempty"`
			BClass string `json:"B-Class,omitempty"`
			BCLASS string `json:"B-CLASS,omitempty"`
			Field5 string `json:",omitempty"`
		} `json:"fares"`
		Company struct {
			Id         string  `json:"id"`
			CompanySKU *string `json:"companySKU"`
			Name       string  `json:"name"`
			TenantId   *string `json:"tenantId"`
		} `json:"company"`
		RouteName      string  `json:"routeName"`
		AvailableSeats int     `json:"availableSeats"`
		StartCounter   string  `json:"startCounter"`
		EndCounter     *string `json:"endCounter"`
		ArrivaleTime   *string `json:"arrivaleTime"`
		DepartureTime  string  `json:"departureTime"`
		BoardingPoints []struct {
			ReportingBranchId int     `json:"reportingBranchId"`
			CounterName       string  `json:"counterName"`
			ReportingDate     *string `json:"reportingDate"`
			ReportingTime     string  `json:"reportingTime"`
			ScheduleTime      string  `json:"scheduleTime"`
		} `json:"boardingPoints"`
		DroppingPoints []struct {
			ReportingBranchId int     `json:"reportingBranchId"`
			CounterName       string  `json:"counterName"`
			ReportingDate     *string `json:"reportingDate"`
			ReportingTime     string  `json:"reportingTime"`
			ScheduleTime      string  `json:"scheduleTime"`
		} `json:"droppingPoints"`
		Seats []struct {
			SeatId       int         `json:"seatId"`
			SeatNo       string      `json:"seatNo"`
			SeatTypeId   string      `json:"seatTypeId"`
			Status       string      `json:"status"`
			ColorCode    *string     `json:"colorCode"`
			Fare         string      `json:"fare"`
			DeckTitle    interface{} `json:"deckTitle"`
			ActualFare   *string     `json:"actualFare"`
			SeatDiscount float64     `json:"seatDiscount"`
			Xaxis        int         `json:"xaxis"`
			Yaxis        int         `json:"yaxis"`
		} `json:"seats"`
		SecondDeckSeats   interface{} `json:"secondDeckSeats"`
		UpperDeck         interface{} `json:"upperDeck"`
		SeatCol           int         `json:"seatCol"`
		SeatRow           int         `json:"seatRow"`
		CoachActiveStatus interface{} `json:"coachActiveStatus"`
		ApiProvider       string      `json:"apiProvider"`
		DiscountDetail    *struct {
			DiscountApplicable bool        `json:"discountApplicable"`
			DiscountType       interface{} `json:"discountType"`
			DiscountAmount     interface{} `json:"discountAmount"`
			DiscountSeats      interface{} `json:"discountSeats"`
		} `json:"discountDetail"`
		TicketToCancel       int     `json:"ticketToCancel"`
		ScDiscountPercentage *string `json:"scDiscountPercentage"`
		ScDiscountFlat       *string `json:"scDiscountFlat"`
		CompanyName          string  `json:"companyName"`
		Obsolete             bool    `json:"obsolete"`
		Fare                 string  `json:"fare"`
	} `json:"data"`
	Version   float64 `json:"version"`
	Timestamp int     `json:"timestamp"`
}

func GetBusInfo() []*AvailableSeat {
	client := &http.Client{}
	var data = strings.NewReader(`{"fromStationId":"5ffb5909b086d27be1901df3","toStationId":"5ffb5909b086d27be1901e10","date":"2024-04-04"}`)
	req, err := http.NewRequest("POST", "https://api.bdtickets.com:20102/v1/coaches/search", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "en-GB,en-US;q=0.9,en;q=0.8")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("origin", "https://new.busbd.com.bd")
	req.Header.Set("referer", "https://new.busbd.com.bd/")
	req.Header.Set("sec-ch-ua", `"Brave";v="123", "Not:A-Brand";v="8", "Chromium";v="123"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "cross-site")
	req.Header.Set("sec-gpc", "1")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	busInfo := &BusInfo{}

	json.Unmarshal(bodyText, busInfo)

	as := []*AvailableSeat{}

	for _, coach := range busInfo.Data {
		ct := coach.CoachType
		dt := coach.DepartureTime
		if (ct == "AC" || ct == "Ac") && (dt == "11:00 PM" || dt == "11:45 PM") {
			//log.Printf("%v", coach.CompanyName)
			//log.Printf("%v", coach.DepartureTime)
			//log.Printf("%v", coach.AvailableSeats)
			//log.Println(".........")
			as = append(as, &AvailableSeat{
				CompanyName:    coach.CompanyName,
				DepartureTime:  coach.DepartureTime,
				AvailableSeats: coach.AvailableSeats,
			})
		}

	}

	//fmt.Printf("%s\n", bodyText)

	return as
}

type AvailableSeat struct {
	CompanyName    string
	DepartureTime  string
	AvailableSeats int
}

func main() {

	t := time.NewTicker(2 * time.Second)

	for _ = range t.C {

		x := GetBusInfo()

		for _, seat := range x {
			log.Println(seat)
			log.Println("......")
		}

	}
}

