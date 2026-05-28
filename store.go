package main

import "github.com/Kotaorie/weather_go/model"

type Store struct {
	stations map[string]model.Station
}

func NewStore() *Store {
	return &Store{stations: make(map[string]model.Station)}
}

func (s *Store) Put(st model.Station) {
	s.stations[st.Id] = st
}

// func (s *Store) Has(id string) bool {}

// func (s *Store) Get(id string) (model.Station, bool) {}

// func (s *Store) Delete(id string) bool {}

func (s *Store) All() []model.Station {
	all := make([]model.Station, 0, len(s.stations))
	for _, st := range s.stations {
		all = append(all, st)
	}
	return all
}
