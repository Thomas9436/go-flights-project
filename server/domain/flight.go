package domain

type Flight struct {
	ID              string `json:"id"`
	BookingID       string `json:"bookingId"`
	Status          string `json:"status"`
	DepartureAirport string `json:"departureAirport"`
	ArrivalAirport   string `json:"arrivalAirport"`
	Price           int    `json:"price"`
	Currency        string `json:"currency"`
}

// DTO renvoyé par ton API
type FlightSummary struct {
	BookingID string `json:"bookingId"`
	Status    string `json:"status"`
}
