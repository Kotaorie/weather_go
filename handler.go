package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Kotaorie/weather_go/model"
)

type App struct{ store *Store }

type TransportJSON struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	CountryCode string `json:"country_code"`
	Altitude    int    `json:"altitude"`
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}

func (a *App) listStations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(a.store.All())
}

func (a *App) getStation(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	st, ok := a.store.Get(id)
	if !ok {
		writeError(w, http.StatusNotFound,
			fmt.Sprintf("station %q introuvable", id))
		return
	}
	writeJSON(w, http.StatusOK, st)
}

func (a *App) createStation(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var dto TransportJSON
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&dto); err != nil {
		writeError(w, 400, "JSON invalide: "+err.Error())
		return
	}
	st := model.Station{
		Id:       dto.Id,
		Name:     dto.Name,
		Country:  dto.CountryCode,
		Altitude: dto.Altitude,
	}
	if st.Id == "" {
		writeError(w, 400, "id manquant")
		return
	}
	if a.store.Has(st.Id) {
		writeError(w, 409, "id "+st.Id+" déjà utilisé")
		return
	}
	a.store.Put(st)

	writeJSON(w, 201, st)
}
