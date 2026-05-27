package main

import (
	"fmt"
	"log"

	"github.com/Kotaorie/weather_go/weatherDataJson"
	"github.com/Kotaorie/weather_go/weatherDataXml"
)

func main() {
	fmt.Println("Welcome to Weather!")

	jsonStations, err := weatherDataJson.LoadFromJSON("./weatherDataJson/weather_data.json")
	if err != nil {
		log.Fatal("JSON error:", err)
	}

	xmlStations, err := weatherDataXML.LoadFromXML("./weatherDataXml/weather_data.xml")
	if err != nil {
		log.Fatal("XML error:", err)
	}

	fmt.Println("JSON stations:", len(jsonStations))
	fmt.Println("XML stations:", len(xmlStations))

	totalJSON := 0
	for _, s := range jsonStations {
		totalJSON += len(s.Observations)
	}

	totalXML := 0
	for _, s := range xmlStations {
		totalXML += len(s.Observations)
	}

	fmt.Println("JSON observations:", totalJSON)
	fmt.Println("XML observations:", totalXML)

	if len(jsonStations) == len(xmlStations) && totalJSON == totalXML {
		fmt.Println("Cohérence : OK")
	} else {
		fmt.Println("Cohérence : FAIL")
	}

	maxStationJson, maxWindJson := MaxWindGust(jsonStations)
	fmt.Println("Max Wind json: ", maxWindJson)
	fmt.Println("Max Wind json: ", maxStationJson)

	maxStation, maxWind := MaxWindGust(xmlStations)
	fmt.Println("Max Wind xml: ", maxWind)
	fmt.Println("Max Wind xml: ", maxStation)
}
