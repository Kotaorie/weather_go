package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Kotaorie/weather_go/model"
)

type App struct{ store *Store }

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

type ErrorResponse struct {
	Error string `json:"error"`
	Code  string `json:"code,omitempty"`
}

func writeError(w http.ResponseWriter, status int, code, msg string) {
	writeJSON(w, status, ErrorResponse{Error: msg, Code: code})
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
		writeError(w, http.StatusNotFound, "STATION NOT FOUND",
			fmt.Sprintf("station %q introuvable", id))
		return
	}
	writeJSON(w, http.StatusOK, st)
}

func (a *App) createStation(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var st model.Station
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&st); err != nil {
		writeError(w, 400, "INVALID JSON", "JSON invalide: "+err.Error())
		return
	}
	if st.Id == "" {
		writeError(w, 400, "ID MISSING", "id manquant")
		return

	}
	if a.store.Has(st.Id) {
		writeError(w, 409, "ID ALREADY TAKEN", "id "+st.Id+" déjà utilisé")
		return
	}
	a.store.Put(st)
	writeJSON(w, 201, st)
}

func (a *App) updateStation(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	id := r.PathValue("id")
	var st model.Station
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&st); err != nil {
		writeError(w, 400, "INVALID JSON", "JSON invalide: "+err.Error())
		return
	}
	if st.Id != "" && st.Id != id {
		writeError(w, 400, "WRONG ID BODY VS URL", "incohérence id body vs URL")
		return
	}
	st.Id = id
	created := !a.store.Has(id)
	a.store.Put(st)
	if created {
		writeJSON(w, 201, st)
		return
	}
	writeJSON(w, 200, st)
}

func (a *App) deleteStation(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if !a.store.Delete(id) {
		writeError(w, http.StatusNotFound, "NOT FOUND",
			fmt.Sprintf("station %q introuvable", id))
		return
	}
	w.WriteHeader(http.StatusNoContent) // 204 — pas de body
}

func (a *App) listObservations(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	st, ok := a.store.Get(id)
	if !ok {
		writeError(w, 404, "NOT FOUND", "station introuvable")
		return
	}
	writeJSON(w, http.StatusOK, st.Observations)
}
