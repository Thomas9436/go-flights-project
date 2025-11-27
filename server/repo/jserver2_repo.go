package repo

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"

    "github.com/spf13/viper"
	"aggregator/model"
)

/*
Example shape:

"flight_to_book": [
  {
    "reference": "B30001",
    "status": "confirmed",
    "traveler": {
      "firstName": "Marie",
      "lastName": "Curie"
    },
    "segments": [
      {
        "flight": {
          "number": "AF276",
          "from": "CDG",
          "to": "HND",
          "depart": "2026-01-01T10:00:00Z",
          "arrive": "2026-01-01T23:00:00Z"
        }
      }
    ],
    "total": {
      "amount": 950.0,
      "currency": "EUR"
    }
  }
]
*/

type jServer2Traveler struct {
    FirstName string `json:"firstName"`
    LastName  string `json:"lastName"`
}

type jServer2SegmentFlight struct {
    Number string `json:"number"`
    From   string `json:"from"`
    To     string `json:"to"`
    Depart string `json:"depart"`
    Arrive string `json:"arrive"`
}

type jServer2Segment struct {
    Flight jServer2SegmentFlight `json:"flight"`
}

type jServer2Total struct {
    Amount   float64 `json:"amount"`
    Currency string  `json:"currency"`
}

type jServer2Raw struct {
    Reference string              `json:"reference"`
    Status    string              `json:"status"`
    Traveler  jServer2Traveler    `json:"traveler"`
    Segments  []jServer2Segment   `json:"segments"`
    Total     jServer2Total       `json:"total"`
}

type JServer2Repo struct {
    baseURL string
    client  *http.Client
}

func NewJServer2Repo() *JServer2Repo {
    url := viper.GetString("J_SERVER2_URL")
    if url == "" {
        url = "http://localhost:4002"
    }
    return &JServer2Repo{
        baseURL: url,
        client:  &http.Client{Timeout: 10 * time.Second},
    }
}

func (r *JServer2Repo) FetchFlights(destination string) ([]model.Flight, error) {
    url := fmt.Sprintf("%s/flight_to_book", r.baseURL) // endpoint assumed
    if destination != "" {
        url = url + "?to=" + destination
    }

    resp, err := r.client.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var raw []jServer2Raw
    if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
        return nil, err
    }

    out := make([]model.Flight, 0, len(raw))
    for _, r := range raw {
        // We pick the first segment as the main flight for normalized model
        if len(r.Segments) == 0 {
            continue
        }
        seg := r.Segments[0].Flight
        dep, _ := time.Parse(time.RFC3339, seg.Depart)
        arr, _ := time.Parse(time.RFC3339, seg.Arrive)
        passenger := r.Traveler.FirstName + " " + r.Traveler.LastName
        out = append(out, model.Flight{
            Reference:        r.Reference,
            Status:           r.Status,
            PassengerName:    passenger,
            FlightNumber:     seg.Number,
            DepartureAirport: seg.From,
            ArrivalAirport:   seg.To,
            DepartureTime:    dep,
            ArrivalTime:      arr,
            Price:            r.Total.Amount,
            Currency:         r.Total.Currency,
            RawSource:        "j-server2",
        })
    }

    return out, nil
}