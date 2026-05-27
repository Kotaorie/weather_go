package main

import (
	"fmt"

	//"github.com/Kotaorie/weather_go/weatherDataJson"
	"github.com/Kotaorie/weather_go/weatherDataXml"
)

func main() {
	fmt.Println("Welcome to Weather!")
	//Stations, err := weatherDataJson.LoadFromJSON("./weatherDataJson/weather_data.json")
	Stations, err := weatherDataXML.LoadFromXML("./weatherDataXml/weather_data.xml")

	if err != nil {
	}
	fmt.Println(len(Stations))
}
