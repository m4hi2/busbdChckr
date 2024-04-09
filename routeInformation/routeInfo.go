package routeInformation

import (
	"errors"
	availableBus "github.com/JahidNishat/BusTicketChecker/busbdChckr/busInformation"
	"github.com/JahidNishat/BusTicketChecker/busbdChckr/routeInformation/models"
	"github.com/JahidNishat/BusTicketChecker/busbdChckr/stations"
	"strings"
)

func GetBusInfo(source, destination, date string) ([]*models.ResponsePld, int, error) {
	sourceID := stations.StationCodeToStationID[strings.ToLower(source)]
	if sourceID == "" {
		return nil, 0, errors.New("invalid source")
	}
	destID := stations.StationCodeToStationID[strings.ToLower(destination)]
	if destID == "" {
		return nil, 0, errors.New("invalid destination")
	}
	data := availableBus.RequestPld{
		FromStationId: sourceID,
		ToStationId:   destID,
		Date:          date,
	}
	bInfo, err := availableBus.GetAvailableBusInformation(data)
	if err != nil {
		return nil, 0, err
	}
	var notificationForUser []*models.ResponsePld
	var availableSeats int
	for _, coach := range bInfo.Data {
		if coach.AvailableSeats > 0 {
			notificationForUser = append(notificationForUser, &models.ResponsePld{
				RouteId:        coach.RouteId,
				CoachNo:        coach.CoachNo,
				RouteName:      coach.RouteName,
				AvailableSeats: coach.AvailableSeats,
				StartCounter:   coach.StartCounter,
				DepartureTime:  coach.DepartureTime,
				CompanyName:    coach.CompanyName,
				Fare:           coach.Fare,
			})
			availableSeats += coach.AvailableSeats
		}
	}
	return notificationForUser, availableSeats, nil
}
