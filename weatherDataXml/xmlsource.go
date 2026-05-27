package weatherDataXML

import (
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/Kotaorie/weather_go/model"
)

type Stations struct {
	Station []Station `xml:"station"`
}

type Station struct {
	Country      string        `xml:"country,attr"`
	Name         string        `xml:"name"`
	Coordinates  Coordinates   `xml:"coordinates"`
	Hardware     Hardware      `xml:"hardware"`
	Observations []Observation `xml:"observations>observation"`
}

type Coordinates struct {
	Lat      float64 `xml:"lat,attr"`
	Lon      float64 `xml:"lon,attr"`
	Altitude int     `xml:"altitude,attr"`
}

type Hardware struct {
	Vendor string `xml:"vendor,attr"`
	Model  string `xml:"model,attr"`
	Since  string `xml:"since,attr"`
}

type Observation struct {
	At       string    `xml:"at,attr"`
	Sky      string    `xml:"sky,attr"`
	Measures []Measure `xml:"measure"`
	Wind     WindXML   `xml:"wind"`
	Note     *string   `xml:"note"`
}

type Measure struct {
	Type  string `xml:"type,attr"`
	Value string `xml:",chardata"`
}

type WindXML struct {
	Speed     float32 `xml:"speed,attr"`
	Direction float32 `xml:"direction,attr"`
}

func powerRangersTransformationXML(input Stations) []model.Station {
	result := make([]model.Station, 0, len(input.Station))

	for _, z := range input.Station {
		if len(z.Observations) == 0 {
			continue
		}
		obs := z.Observations[0]
		t, err := time.Parse("2006-01-02", z.Hardware.Since)
		if err != nil {
			fmt.Println("Error parsing date:", err)
		}
		var temperature float32
		var condition string

		for _, m := range obs.Measures {
			switch m.Type {
			case "temperature":
				v, err := strconv.ParseFloat(m.Value, 32)
				if err == nil {
					temperature = float32(v)
				}
			}
		}
		condition = obs.Sky
		result = append(result, model.Station{
			Country: z.Country,
			Location: model.Location{
				Longitude: z.Coordinates.Lon,
				Latitude:  z.Coordinates.Lat,
			},
			Altitude: z.Coordinates.Altitude,
			Device: model.Device{
				Type:         z.Hardware.Model,
				Manufacturer: z.Hardware.Vendor,
				InstalledOn:  t,
			},
			Temperature: temperature,
			Condition:   condition,
			Wind: model.Wind{
				Speed: z.Observations[0].Wind.Speed,
				Deg:   z.Observations[0].Wind.Direction,
			},
			Note: z.Observations[0].Note,
		})
	}

	return result
}

func LoadFromXML(path string) ([]model.Station, error) {
	var stationsXML Stations

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading xml file: %v", err)
	}
	err = xml.Unmarshal(data, &stationsXML)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling xml: %v", err)
	}
	final := powerRangersTransformationXML(stationsXML)
	fmt.Println(final[1])
	return final, nil
}
