package availableBusInformation

import (
	"encoding/json"
	"github.com/fahimimam/busbdChckr/businfo/client"
	"github.com/fahimimam/busbdChckr/businfo/models"
	"net/http"
)

type RequestPld struct {
	Date          string `json:"date"`
	Identifier    string `json:"identifier"`
	StructureType string `json:"structureType"`
}

func GetAvailableBusInformation(data RequestPld) (*models.BusInfo, error) {
	buf, err := BodyBuffer(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://api.bdtickets.com:20102/v1/coaches/search", buf)
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

	busInfo := &models.BusInfo{}
	err = json.NewDecoder(resp.Body).Decode(busInfo)
	if err != nil {
		return nil, err
	}

	return busInfo, nil
}
