package main

import (
	"fmt"

	"github.com/Kotaorie/weather_go/model"
)

func FilterByCountry(stations []model.Station, iso string) []model.Station {
	result := make([]model.Station, 0)

	for _, s := range stations {
		if s.Country == iso {
			result = append(result, s)
		}
	}

	return result
}

func AvgTemperature(s model.Station) float64 {
	if len(s.Observations) == 0 {
		return 0
	}

	var sum float64
	var count float64

	for _, o := range s.Observations {
		sum += float64(o.Temperature)
		count++
	}

	return sum / count
}

func MaxWindGust(stations []model.Station) (model.Station, float64) {
	var maxStation model.Station
	var max float64

	for _, s := range stations {
		for _, o := range s.Observations {
			if float64(o.Wind.Speed) > max {
				max = float64(o.Wind.Speed)
				maxStation = s
			}
		}
	}
	fmt.Println(maxStation)
	return maxStation, max
}

func CountByCountry(stations []model.Station) map[string]int {
	result := make(map[string]int)
	for _, s := range stations {
		result[s.Country]++
	}
	return result
}
