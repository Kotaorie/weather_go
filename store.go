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

func (s *Store) Has(id string) bool {
	_, ok := s.stations[id]
	return ok
}

func (s *Store) Get(id string) (model.Station, bool) {
	for _, st := range s.stations {
		if st.Id == id {
			return st, true
		}
	}
	return model.Station{}, false
}

func (s *Store) All() []model.Station {
	all := make([]model.Station, 0, len(s.stations))
	for _, st := range s.stations {
		all = append(all, st)
	}
	return all
}

func (s *Store) Delete(id string) bool {
	if _, ok := s.stations[id]; !ok {
		return false
	}
	delete(s.stations, id)
	return true
}
