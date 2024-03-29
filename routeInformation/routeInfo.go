package routeInformation

import (
	availableBus "github.com/fahimimam/busbdChckr/routeInformation/busInformation"
	"github.com/fahimimam/busbdChckr/routeInformation/models"
)

func BuildUrl(path string) string {
	return BDTicketHost + BDTicketPort + path
}

func GetBusInfo(data availableBus.RequestPld) ([]*models.NotificationPld, error) {
	bInfo, err := availableBus.GetAvailableBusInformation(data)
	if err != nil {
		return nil, err
	}
	var notificationForUser []*models.NotificationPld
	for _, coach := range bInfo.Data {
		notificationForUser = append(notificationForUser, &models.NotificationPld{
			RouteId:        coach.RouteId,
			CoachNo:        coach.CoachNo,
			RouteName:      coach.RouteName,
			AvailableSeats: coach.AvailableSeats,
			StartCounter:   coach.StartCounter,
			ArrivaleTime:   coach.ArrivaleTime,
			DepartureTime:  coach.DepartureTime,
			CompanyName:    coach.CompanyName,
			Fare:           coach.Fare,
		})
	}
	return notificationForUser, nil
}
