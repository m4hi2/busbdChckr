package busInformation

import (
	"encoding/json"
	"github.com/m4hi2/busbdChckr/routeInformation/client"
	"net/http"
)

const (
	Path = "/v1/coaches/search"
)

type RequestPld struct {
	FromStationId string `json:"fromStationId"`
	ToStationId   string `json:"toStationId"`
	Date          string `json:"date"`
}

type BusInfo struct {
	Message interface{} `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    []struct {
		Id          string   `json:"id"`
		Date        string   `json:"date"`
		RouteId     string   `json:"routeId"`
		CoachNo     string   `json:"coachNo"`
		CoachType   string   `json:"coachType"`
		SeatTypesId []string `json:"seatTypesId"`
		Fares       struct {
			EClass string `json:"E-Class,omitempty"`
			ECLASS string `json:"E-CLASS,omitempty"`
			BClass string `json:"B-Class,omitempty"`
			BCLASS string `json:"B-CLASS,omitempty"`
		} `json:"fares"`
		RouteName      string  `json:"routeName"`
		AvailableSeats int     `json:"availableSeats"`
		StartCounter   string  `json:"startCounter"`
		EndCounter     *string `json:"endCounter"`
		ArrivaleTime   *string `json:"arrivaleTime"`
		DepartureTime  string  `json:"departureTime"`
		DiscountDetail *struct {
			DiscountApplicable bool        `json:"discountApplicable"`
			DiscountType       interface{} `json:"discountType"`
			DiscountAmount     interface{} `json:"discountAmount"`
			DiscountSeats      interface{} `json:"discountSeats"`
		} `json:"discountDetail"`
		CompanyName string `json:"companyName"`
		Fare        string `json:"fare"`
	} `json:"data"`
}

func GetAvailableBusInformation(data RequestPld) (*BusInfo, error) {
	buf, err := BodyBuffer(data)
	if err != nil {
		return nil, err
	}
	url := BuildUrl(Path)
	req, err := http.NewRequest(http.MethodPost, url, buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("content-type", "application/json")

	resp, err := client.GetClient().Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	busInfo := &BusInfo{}
	err = json.NewDecoder(resp.Body).Decode(busInfo)
	if err != nil {
		return nil, err
	}

	return busInfo, nil
}
