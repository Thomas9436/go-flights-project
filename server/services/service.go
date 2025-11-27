package service

import (
	"aggregator/model"
	"aggregator/repo"
	"aggregator/sorters"
	"errors"
)

// FlightService merges repos and exposes business logic.
type FlightService struct {
	repos []repo.FlightRepository
}

func NewFlightService(repos ...repo.FlightRepository) *FlightService {
	return &FlightService{repos: repos}
}

func (s *FlightService) FetchAndMerge(destination string, sortKey string) ([]model.Flight, error) {
	if len(s.repos) == 0 {
		return nil, errors.New("no repositories configured")
	}
	merged := make([]model.Flight, 0)
	for _, r := range s.repos {
		f, err := r.FetchFlights(destination)
		if err != nil {
			return nil, err
		}
		merged = append(merged, f...)
	}

	sorters.SortByKey(merged, sortKey)

	return merged, nil
}
