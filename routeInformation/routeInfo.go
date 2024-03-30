package routeInformation

import (
	availableBus "github.com/JahidNishat/BusTicketChecker/busbdChckr/busInformation"
	"github.com/JahidNishat/BusTicketChecker/busbdChckr/routeInformation/models"
	"github.com/JahidNishat/BusTicketChecker/busbdChckr/stations"
	"strings"
)

func GetBusInfo(source, destination, date string) ([]*models.ResponsePld, error) {
	sourceID := stations.StationCodeToStationID[strings.ToLower(source)]
	destID := stations.StationCodeToStationID[strings.ToLower(destination)]
	data := availableBus.RequestPld{
		FromStationId: sourceID,
		ToStationId:   destID,
		Date:          date,
	}
	bInfo, err := availableBus.GetAvailableBusInformation(data)
	if err != nil {
		return nil, err
	}
	var notificationForUser []*models.ResponsePld
	for _, coach := range bInfo.Data {
		notificationForUser = append(notificationForUser, &models.ResponsePld{
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
