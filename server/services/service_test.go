package service

import (
	"errors"
	"testing"
	"time"

	"aggregator/model"

	"github.com/stretchr/testify/assert"
)

// mock repo
type mockRepo struct {
	flights []model.Flight
	err     error
}

func (m *mockRepo) FetchFlights(dest string) ([]model.Flight, error) {
	if m.err != nil {
		return nil, m.err
	}
	if dest == "" {
		return m.flights, nil
	}
	out := []model.Flight{}
	for _, f := range m.flights {
		if f.ArrivalAirport == dest || f.DepartureAirport == dest {
			out = append(out, f)
		}
	}
	return out, nil
}

func mf(price float64, dep, arr, arrAirport string) model.Flight {
	d, _ := time.Parse(time.RFC3339, dep)
	a, _ := time.Parse(time.RFC3339, arr)
	return model.Flight{Price: price, DepartureTime: d, ArrivalTime: a, ArrivalAirport: arrAirport}
}

func TestFetchAndMerge_SortsByPrice(t *testing.T) {
	r1 := &mockRepo{flights: []model.Flight{
		mf(900.0, "2026-01-01T10:00:00Z", "2026-01-01T14:00:00Z", "HND"),
	}}
	r2 := &mockRepo{flights: []model.Flight{
		mf(500.0, "2026-01-01T08:00:00Z", "2026-01-01T12:00:00Z", "HND"),
	}}
	s := NewFlightService(r1, r2)
	out, err := s.FetchAndMerge("HND", "price")
	assert.NoError(t, err)
	assert.Len(t, out, 2)
	assert.Equal(t, 500.0, out[0].Price)
	assert.Equal(t, 900.0, out[1].Price)
}

func TestFetchAndMerge_RepoError(t *testing.T) {
	rErr := &mockRepo{err: errors.New("boom")}
	s := NewFlightService(rErr)
	_, err := s.FetchAndMerge("", "price")
	assert.Error(t, err)
}
