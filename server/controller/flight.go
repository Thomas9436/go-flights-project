package controller

import (
	"aggregator/services"
	"encoding/json"
	"net/http"
)

type FlightsController struct {
	service *services.FlightService
}

func NewFlightsController(service *services.FlightService) *FlightsController {
	return &FlightsController{service: service}
}

func (c *FlightsController) Flight(w http.ResponseWriter, r *http.Request) {
	flights, err := c.service.List(r.Context())
	if err != nil {
		http.Error(w, "upstream fetch failed: "+err.Error(), http.StatusBadGateway)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_ = json.NewEncoder(w).Encode(flights)
}
