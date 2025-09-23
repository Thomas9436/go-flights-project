package config

import (
	"aggregator/controller"
	"aggregator/repos"
	"aggregator/services"
	"net/http"
)

func NewRouter(cfg Config) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", controller.Health)

	// DI: repo -> service -> controller
	flightRepo := repos.NewHTTPFlightRepo(cfg.J1BaseURL)
	flightSvc  := services.NewFlightService(flightRepo)
	flightCtrl := controller.NewFlightsController(flightSvc)

	mux.HandleFunc("/flight", flightCtrl.Flight)
	return mux
}
