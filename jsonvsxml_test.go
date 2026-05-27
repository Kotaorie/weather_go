package main

import (
	"testing"

	"github.com/Kotaorie/weather_go/weatherDataJson"
	"github.com/Kotaorie/weather_go/weatherDataXml"
)

func TestJSONVsXML(t *testing.T) {

	jsonStations, err := weatherDataJson.LoadFromJSON("./weatherDataJson/weather_data.json")
	if err != nil {
		t.Fatal(err)
	}

	xmlStations, err := weatherDataXML.LoadFromXML("./weatherDataXml/weather_data.xml")
	if err != nil {
		t.Fatal(err)
	}

	if len(jsonStations) != len(xmlStations) {
		t.Fatalf("stations mismatch: json=%d xml=%d", len(jsonStations), len(xmlStations))
	}

	jsonObs := 0
	xmlObs := 0

	for _, s := range jsonStations {
		jsonObs += len(s.Observations)
	}

	for _, s := range xmlStations {
		xmlObs += len(s.Observations)
	}

	if jsonObs != xmlObs {
		t.Fatalf("observations mismatch: json=%d xml=%d", jsonObs, xmlObs)
	}
}
