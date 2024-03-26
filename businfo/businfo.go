package businfo

import (
	businfo "github.com/fahimimam/busbdChckr/businfo/availableBusInformation"
	"github.com/fahimimam/busbdChckr/businfo/models"
)

func GetBusInfoV2(data businfo.RequestPld) ([]*models.NotificationPld, error) {
	bInfo, err := businfo.GetAvailableBusInformation(data)
	if err != nil {
		return nil, err
	}
	var as []*models.NotificationPld
	for _, coach := range bInfo.Data {
		as = append(as, &models.NotificationPld{
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
	return as, nil
}
