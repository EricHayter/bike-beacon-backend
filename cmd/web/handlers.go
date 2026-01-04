package main

import (
	"fmt"
	"github.com/erichayter/bike-beacon-backend/internal/models"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

func (app *application) getStation(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		app.errorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	// query from the model otherwise
	station, err := app.repairStations.Get(r.Context(), id)
	if err != nil {
		app.errorJSON(w, http.StatusNotFound, err.Error())
		return
	}

	app.writeJSON(w, http.StatusOK, station, http.Header{})
}

func (app *application) getStations(w http.ResponseWriter, r *http.Request) {
	lngStr := r.URL.Query().Get("lng")
	if lngStr == "" {
		app.errorJSON(w, http.StatusBadRequest, "Missing \"lng\" in query parameters")
		return
	}
	lng, err := strconv.ParseFloat(lngStr, 64)
	if err != nil {
		app.errorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	latStr := r.URL.Query().Get("lat")
	if latStr == "" {
		app.errorJSON(w, http.StatusBadRequest, "Missing \"lat\" in query parameters")
		return
	}
	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		app.errorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	location := models.Point{Lng: lng, Lat: lat}

	stations, err := app.repairStations.GetNearby(r.Context(), location)
	if err != nil {
		app.errorJSON(w, http.StatusNotFound, err.Error())
		return
	}

	app.writeJSON(w, http.StatusOK, stations, http.Header{})
}

func (app *application) getTools(w http.ResponseWriter, r *http.Request) {
	stationId, err := uuid.Parse(r.PathValue("id"))
	tools, err := app.tools.Get(r.Context(), stationId)
	if err != nil {
		app.errorJSON(w, http.StatusNotFound, err.Error())
		return
	}

	app.writeJSON(w, http.StatusOK, tools, http.Header{})
}

func (app *application) getPhotos(w http.ResponseWriter, r *http.Request) {
	stationId, err := uuid.Parse(r.PathValue("id"))
	photos, err := app.repairStationPhotos.Get(r.Context(), stationId)
	if err != nil {
		app.errorJSON(w, http.StatusNotFound, err.Error())
		return
	}
	var photoUrls []string
	for _, photo := range photos {
		// TODO this logically should probably be reading from a configuration
		// or .env file. Fix this once done so that it's not hardcoded
		photoUrl := fmt.Sprintf("http://192.168.0.138:9000/images/%s", photo.PhotoKey)
		photoUrls = append(photoUrls, photoUrl)
	}

	app.writeJSON(w, http.StatusOK, photoUrls, http.Header{})
}
