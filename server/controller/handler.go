package controller

import (
	service "aggregator/services"
	"encoding/json"
	"fmt"
	"net/http"
)

// Handler wraps the service
type Handler struct {
	Svc *service.FlightService
}

func NewHandler(svc *service.FlightService) *Handler {
	return &Handler{Svc: svc}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/health", h.Health)
	mux.HandleFunc("/flight", h.Flights)
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	// Per your request — return 201 Created (StatusCreated)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "OK")
}

func (h *Handler) Flights(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	destination := r.URL.Query().Get("to")
	sortKey := r.URL.Query().Get("sort_by")
	if sortKey == "" {
		sortKey = "price"
	}

	flights, err := h.Svc.FetchAndMerge(destination, sortKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(flights); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
