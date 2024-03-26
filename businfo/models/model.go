package models

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

type NotificationPld struct {
	RouteId        string  `json:"routeId"`
	CoachNo        string  `json:"coachNo"`
	RouteName      string  `json:"routeName"`
	AvailableSeats int     `json:"availableSeats"`
	StartCounter   string  `json:"startCounter"`
	ArrivaleTime   *string `json:"arrivaleTime"`
	DepartureTime  string  `json:"departureTime"`
	CompanyName    string  `json:"companyName"`
	Fare           string  `json:"fare"`
}
