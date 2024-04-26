package notifier

import (
	"bytes"
	"crypto/sha1"
	_ "embed"
	"encoding/hex"
	"encoding/json"
	"fmt"
	rm "github.com/m4hi2/busbdChckr/routeInformation/models"
	"log"
	"reflect"
	"text/template"
)

//go:embed "response.tmpl"
var responseTmpl string

var result *template.Template

func init() {
	var err error
	if result == nil {
		result, err = template.New("responseTmpl").Parse(responseTmpl)
		if err != nil {
			log.Printf("Error parsing template: %v", err)
			return
		}
	}
}

type ResponseData struct {
	CoachNo        string `json:"coachNo"`
	RouteName      string `json:"routeName"`
	AvailableSeats int    `json:"availableSeats"`
	StartCounter   string `json:"startCounter"`
	DepartureTime  string `json:"departureTime"`
	CompanyName    string `json:"companyName"`
	Fare           string `json:"fare"`
}

func ResponseStruct(busInfo *rm.ResponsePld) string {
	responseData := &ResponseData{
		CoachNo:        busInfo.CoachNo,
		RouteName:      busInfo.RouteName,
		AvailableSeats: busInfo.AvailableSeats,
		StartCounter:   busInfo.StartCounter,
		DepartureTime:  busInfo.DepartureTime,
		CompanyName:    busInfo.CompanyName,
		Fare:           busInfo.Fare,
	}

	var b bytes.Buffer
	err := result.Execute(&b, responseData)
	if err != nil {
		log.Printf("Error executing template: %v", err)
	}

	return b.String()
}

func StringifyStruct(data interface{}) string {
	val := reflect.ValueOf(data)
	typ := reflect.TypeOf(data)

	if val.Kind() != reflect.Struct {
		return "Not a struct"
	}

	var str string
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		str += fmt.Sprintf("%s: %v\n", fieldType.Name, field.Interface())
	}

	return str
}

type HashedDataPld struct {
	RouteId     string `json:"routeId"`
	CoachNo     string `json:"coachNo"`
	RouteName   string `json:"routeName"`
	CompanyName string `json:"companyName"`
}

func HashData(busInfo *rm.ResponsePld) string {
	hashedDataPld := &HashedDataPld{
		RouteId:     busInfo.RouteId,
		CoachNo:     busInfo.CoachNo,
		RouteName:   busInfo.RouteName,
		CompanyName: busInfo.CompanyName,
	}

	jsonData, _ := json.Marshal(hashedDataPld)

	sha1Hash := sha1.New()
	sha1Hash.Write(jsonData)
	hashBytes := sha1Hash.Sum(nil)
	sha1String := hex.EncodeToString(hashBytes)

	return sha1String
}
