package model

import "time"

type Station struct {
	Id           string
	Name         string
	Country      string
	Location     Location
	Altitude     int
	Device       Device
	Observations []Observation
}

type Observation struct {
	Temperature float32
	Condition   string
	Wind        Wind
	Note        *string
}

type Device struct {
	Type         string
	Manufacturer string
	InstalledOn  time.Time
}
type Location struct {
	Longitude float64
	Latitude  float64
}

type Wind struct {
	Speed float32 //KMH
	Deg   float32
}
