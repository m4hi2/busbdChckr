package models

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
