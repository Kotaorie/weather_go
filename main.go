package main

import (
	"fmt"

	"github.com/Kotaorie/weather_go/weatherDataJson"
)

func main() {
	fmt.Println("Welcome to Weather!")
	Stations, err := weatherDataJson.LoadFromJSON("./weatherDataJson/weather_data.json")
	if err != nil {
	}
	fmt.Println(len(Stations))
}
