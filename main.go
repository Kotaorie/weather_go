package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Kotaorie/weather_go/model"

	//import j3 tp
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})
	http.ListenAndServe(":8080", mux)

	startJSON := time.Now()
	jsonStations, err := LoadFromAuto("./weatherDataJson/weather_data.json")
	if err != nil {
		log.Fatal("JSON error:", err)
	}
	jsonDuration := time.Since(startJSON)

	startXML := time.Now()
	xmlStations, err := LoadFromAuto("./weatherDataXml/weather_data.xml")
	if err != nil {
		log.Fatal("XML error:", err)
	}
	xmlDuration := time.Since(startXML)

	fmt.Printf("JSON Duration: %s vs XML Duration %s\n", jsonDuration, xmlDuration)
	fmt.Printf("JSON stations: %d and XML stations : %d\n", len(jsonStations), len(xmlStations))

	totalJSON := 0
	for _, s := range jsonStations {
		totalJSON += len(s.Observations)
	}

	totalXML := 0
	for _, s := range xmlStations {
		totalXML += len(s.Observations)
	}

	fmt.Printf("JSON observations: %d and XML observations: %d\n", totalJSON, totalXML)

	if len(jsonStations) == len(xmlStations) && totalJSON == totalXML {
		fmt.Println("Cohérence : OK")
	} else {
		fmt.Println("Cohérence : FAIL")
	}

	allStations := []model.Station{}

	allStations = append(allStations, jsonStations...)
	allStations = append(allStations, xmlStations...)

	maxStation, maxWind := MaxWindGust(allStations)
	fmt.Printf("Max Wind xml: %f for %s\n", maxWind, maxStation.Id)

	bordeauxStation, true := FindStationByID(allStations, "FR-BOR-001")
	if true {
		avgTemp := AvgTemperature(bordeauxStation)
		fmt.Println("bordeaux Station: ", bordeauxStation)
		fmt.Println("avg temp: ", avgTemp)
	}

	franceStation := FilterByCountry(allStations, "FR")
	fmt.Println("france Station: ", len(franceStation))

}
