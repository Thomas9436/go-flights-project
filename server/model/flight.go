package model

import "time"

// Flight is the normalized model used by the service & API.
type Flight struct {
	// Primary IDs (one of these may be set depending on source)
	BookingID string `json:"bookingId,omitempty"`
	Reference string `json:"reference,omitempty"`

	// Common fields
	Status           string    `json:"status,omitempty"`
	PassengerName    string    `json:"passengerName,omitempty"`
	FlightNumber     string    `json:"flightNumber,omitempty"`
	DepartureAirport string    `json:"departureAirport,omitempty"`
	ArrivalAirport   string    `json:"arrivalAirport,omitempty"`
	DepartureTime    time.Time `json:"departureTime,omitempty"`
	ArrivalTime      time.Time `json:"arrivalTime,omitempty"`
	Price            float64   `json:"price,omitempty"`
	Currency         string    `json:"currency,omitempty"`

	RawSource string `json:"rawSource,omitempty"` // e.g. j-server1 or j-server2
}
