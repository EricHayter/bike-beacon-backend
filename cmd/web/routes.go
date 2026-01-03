package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /stations/{id}", app.getStation)
	mux.HandleFunc("GET /stations/{id}/tools", app.getTools)
	mux.HandleFunc("GET /stations/{id}/photos", app.getPhotos)
	mux.HandleFunc("GET /stations", app.getStations)

	return mux
}
