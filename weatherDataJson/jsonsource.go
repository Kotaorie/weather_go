package weatherDataJson

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/Kotaorie/weather_go/model"
)

type Stations struct {
	Station []Station `json:"stations"`
}

type Station struct {
	Id          string        `json:"id"`
	Name        string        `json:"name"`
	Country     string        `json:"country"`
	Location    Location      `json:"location"`
	Altitude    int           `json:"altitude_m"`
	Device      Device        `json:"device"`
	Observation []Observation `json:"observations"`
}

type Observation struct {
	Temperature float32 `json:"temperature_celsius"`
	Condition   string  `json:"conditions"`
	Wind        Wind    `json:"wind"`
	Note        *string `json:"note"`
}

type Device struct {
	Type         string `json:"type"`
	Manufacturer string `json:"manufacturer"`
	InstalledOn  string `json:"installed_on"`
}
type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type Wind struct {
	Speed float32 `json:"speed_kmh"` //KMH
	Deg   float32 `json:"direction_deg"`
}

func countryToISO(country string) [2]rune {
	switch country {
	case "France":
		return [2]rune{'F', 'R'}
	case "Espagne":
		return [2]rune{'E', 'S'}
	case "Belgique":
		return [2]rune{'B', 'L'}
	case "Portugal":
		return [2]rune{'P', 'T'}
	case "Italie":
		return [2]rune{'I', 'L'}
	case "Allemagne":
		return [2]rune{'A', 'L'}
	case "Pays-Bas":
		return [2]rune{'P', 'T'}
	case "Autriche":
		return [2]rune{'A', 'U'}
	case "Suisse":
		return [2]rune{'S', 'U'}
	case "Danemark":
		return [2]rune{'D', 'K'}
	case "Suède":
		return [2]rune{'S', 'E'}
	case "Norvège":
		return [2]rune{'N', 'O'}
	case "Pologne":
		return [2]rune{'P', 'T'}
	case "Tchéquie":
		return [2]rune{'C', 'Z'}
	}
	return [2]rune{}
}

func powerRangersTransformation(input Stations) []model.Station {
	result := make([]model.Station, 0, len(input.Station))
	for _, z := range input.Station {
		t, _ := time.Parse("2006-01-02", z.Device.InstalledOn)
		iso := countryToISO(z.Country)
		obs := make([]model.Observation, 0, len(z.Observation))
		for _, o := range z.Observation {
			obs = append(obs, model.Observation{
				Temperature: o.Temperature,
				Condition:   o.Condition,
				Wind:        model.Wind(o.Wind),
				Note:        o.Note,
			})
		}
		result = append(result, model.Station{
			Id:      z.Id,
			Country: string(iso[0]) + string(iso[1]),
			Location: model.Location{
				Longitude: z.Location.Longitude,
				Latitude:  z.Location.Latitude,
			},
			Altitude: z.Altitude,
			Device: model.Device{
				Type:         z.Device.Type,
				Manufacturer: z.Device.Manufacturer,
				InstalledOn:  t,
			},
			Observations: obs,
		})
	}
	return result
}

func LoadFromJSON(path string) ([]model.Station, error) {
	var StationsJson Stations
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("An error occurred while reading json file:  %v", err.Error())
	}
	err2 := json.Unmarshal([]byte(data), &StationsJson)
	if err2 != nil {
		fmt.Println(err2)
		return nil, fmt.Errorf("An error occurred while Unmarshal the json :  %v", err2.Error())
	}
	finalJSON := powerRangersTransformation(StationsJson)
	return finalJSON, err
}
