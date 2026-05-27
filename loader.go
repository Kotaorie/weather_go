package main

import (
	"fmt"
	"path/filepath"

	"github.com/Kotaorie/weather_go/model"
	"github.com/Kotaorie/weather_go/weatherDataJson"
	"github.com/Kotaorie/weather_go/weatherDataXml"
)

func LoadFromAuto(path string) ([]model.Station, error) {
	extension := filepath.Ext(path)
	switch extension {
	case ".json":
		return weatherDataJson.LoadFromJSON(path)
	case ".xml":
		return weatherDataXML.LoadFromXML(path)
	default:
		return nil, fmt.Errorf("unsupported file formats: %s", extension)
	}
}
