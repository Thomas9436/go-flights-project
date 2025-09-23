package services

import (
	"aggregator/domain"
	"aggregator/repos"
	"context"
)

type FlightService struct {
	repo repos.FlightRepository
}

func NewFlightService(r repos.FlightRepository) *FlightService {
	return &FlightService{repo: r}
}

func (s *FlightService) List(ctx context.Context) ([]domain.Flight, error) {
	return s.repo.List(ctx)
}

// Transforme la liste complète en résumés
func (s *FlightService) Summaries(ctx context.Context) ([]domain.FlightSummary, error) {
	flights, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	out := make([]domain.FlightSummary, 0, len(flights))
	for _, f := range flights {
		out = append(out, domain.FlightSummary{BookingID: f.BookingID, Status: f.Status})
	}
	return out, nil
}
