package businfo

import (
	"github.com/m4hi2/busbdChckr/businfo/availableBusInformation"
	"github.com/m4hi2/busbdChckr/businfo/models"
)

type RequestPld struct {
	Date          string `json:"date"`
	Identifier    string `json:"identifier"`
	StructureType string `json:"structureType"`
}

func GetBusInfoV2(data RequestPld) ([]*models.NotificationPld, error) {
	bInfo, err := availableBusInformation.GetAvailableBusInformation(data)
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
