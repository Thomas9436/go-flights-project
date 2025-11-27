package repo

import "aggregator/model"

// FlightRepository abstracts a flight source.
type FlightRepository interface {
	// FetchFlights returns flights for a destination (arrival airport) or all if empty
	FetchFlights(destination string) ([]model.Flight, error)
}
