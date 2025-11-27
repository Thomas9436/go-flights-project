package repo

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"

    "github.com/spf13/viper"
	"aggregator/model"
)

type jServer1Raw struct {
    BookingID      string  `json:"bookingId"`
    Status         string  `json:"status"`
    PassengerName  string  `json:"passengerName"`
    FlightNumber   string  `json:"flightNumber"`
    DepartureAirport string `json:"departureAirport"`
    ArrivalAirport   string `json:"arrivalAirport"`
    DepartureTime  string  `json:"departureTime"` // ISO8601
    ArrivalTime    string  `json:"arrivalTime"`
    Price          float64 `json:"price"`
    Currency       string  `json:"currency"`
}

type JServer1Repo struct {
    baseURL string
    client  *http.Client
}

func NewJServer1Repo() *JServer1Repo {
    url := viper.GetString("J_SERVER1_URL")
    if url == "" {
        url = "http://localhost:4001"
    }
    return &JServer1Repo{
        baseURL: url,
        client:  &http.Client{Timeout: 10 * time.Second},
    }
}

func (r *JServer1Repo) FetchFlights(destination string) ([]model.Flight, error) {
    url := fmt.Sprintf("%s/flights", r.baseURL)
    if destination != "" {
        url = url + "?to=" + destination
    }

    resp, err := r.client.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var raw []jServer1Raw
    if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
        return nil, err
    }

    out := make([]model.Flight, 0, len(raw))
    for _, r := range raw {
        dep, _ := time.Parse(time.RFC3339, r.DepartureTime)
        arr, _ := time.Parse(time.RFC3339, r.ArrivalTime)
        out = append(out, model.Flight{
            BookingID:       r.BookingID,
            Status:          r.Status,
            PassengerName:   r.PassengerName,
            FlightNumber:    r.FlightNumber,
            DepartureAirport: r.DepartureAirport,
            ArrivalAirport:   r.ArrivalAirport,
            DepartureTime:   dep,
            ArrivalTime:     arr,
            Price:           r.Price,
            Currency:        r.Currency,
            RawSource:       "j-server1",
        })
    }
    return out, nil
}