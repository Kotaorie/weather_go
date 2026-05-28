package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Kotaorie/weather_go/weatherDataJson"
)

func main() {

	//TP3 -------------------------------------------------------------------------------------------
	store := NewStore()
	stations, err := weatherDataJson.LoadFromJSON("./weatherDataJson/weather_data.json")
	if err != nil {
		log.Fatal(err)
	}
	for _, s := range stations {
		store.Put(s)
	}

	app := &App{store: store}
	mux := http.NewServeMux()
	mux.HandleFunc("POST /stations", app.createStation)
	mux.HandleFunc("GET /stations/{id}/observations", app.listObservations)
	mux.HandleFunc("PUT /stations/{id}", app.updateStation)
	mux.HandleFunc("DELETE /stations/{id}", app.deleteStation)
	mux.HandleFunc("GET /stations/{id}", app.getStation)
	mux.HandleFunc("GET /stations", app.listStations)
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})
	http.ListenAndServe(":8080", mux)
	fmt.Println("Listening on port 8080")

	log.Printf("bootstrap : %d stations chargées", len(stations))

	//TP2 -----------------------------------------------------------------------------------------------

	/* startJSON := time.Now()
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
	*/
}
