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
	Country     string   `json:"country"`
	Location    Location `json:"location"`
	Altitude    string   `json:"altitude_m"`
	Device      Device   `json:"device"`
	Temperature float32  `json:"observation:temperature_celsius"`
	Condition   string   `json:"observation:condition"`
	Wind        Wind     `json:"observation:wind"`
	Note        *string  `json:"observation:note"`
}

type Device struct {
	Type         string    `json:"type"`
	Manufacturer string    `json:"manufacturer"`
	InstalledOn  time.Time `json:"installed_on"`
}
type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type Wind struct {
	Speed float32 `json:"speed_kmh"` //KMH
	Deg   float32 `json:"direction_deg"`
}

func powerRangersTransformation(input Stations) []model.Station {
	result := make([]model.Station, 0, len(input.Station))

	for _, s := range input.Station {
		result = append(result, model.Station{
			Country: s.Country,
			Location: model.Location{
				Longitude: s.Location.Longitude,
				Latitude:  s.Location.Latitude,
			},
			Altitude: s.Altitude,
			Device: model.Device{
				Type:         s.Device.Type,
				Manufacturer: s.Device.Manufacturer,
				InstalledOn:  s.Device.InstalledOn,
			},
			Temperature: s.Temperature,
			Condition:   s.Condition,
			Wind: model.Wind{
				Speed: s.Wind.Speed,
				Deg:   s.Wind.Deg,
			},
			Note: s.Note,
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
		return nil, fmt.Errorf("An error occurred while Unmarshal the json :  %v", err2.Error())
	}
	finalJSON := powerRangersTransformation(StationsJson)
	return finalJSON, err
}
