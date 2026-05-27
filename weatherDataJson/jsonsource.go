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
	Country     string        `json:"country"`
	Location    Location      `json:"location"`
	Altitude    int           `json:"altitude_m"`
	Device      Device        `json:"device"`
	Observation []Observation `json:"observations"`
}

type Observation struct {
	Temperature float32 `json:"temperature_celsius"`
	Condition   string  `json:"condition"`
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
		t, err := time.Parse("2006-01-02", z.Device.InstalledOn)
		iso := countryToISO(z.Country)
		if err != nil {
			fmt.Println("Error parsing installed time")
		}
		result = append(result, model.Station{
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
			Temperature: z.Observation[0].Temperature,
			Condition:   z.Observation[0].Condition,
			Wind: model.Wind{
				Speed: z.Observation[0].Wind.Speed,
				Deg:   z.Observation[0].Wind.Deg,
			},
			Note: z.Observation[0].Note,
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
	fmt.Println(finalJSON[10])
	return finalJSON, err
}
