package repos

import (
	"aggregator/domain"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type FlightRepository interface {
	List(ctx context.Context) ([]domain.Flight, error)
}

type HTTPFlightRepo struct {
	BaseURL string
	Client  *http.Client
}

func NewHTTPFlightRepo(baseURL string) *HTTPFlightRepo {
	return &HTTPFlightRepo{
		BaseURL: baseURL,
		Client:  &http.Client{Timeout: 10 * time.Second},
	}
}

func (r *HTTPFlightRepo) List(ctx context.Context) ([]domain.Flight, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/flights", r.BaseURL), nil)
	if err != nil {
		return nil, fmt.Errorf("build req: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := r.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("fetch upstream: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(io.LimitReader(resp.Body, 512))
		return nil, fmt.Errorf("upstream %s: %s", resp.Status, string(b))
	}

	var flights []domain.Flight
	if err := json.NewDecoder(resp.Body).Decode(&flights); err != nil {
		return nil, fmt.Errorf("decode upstream: %w", err)
	}
	return flights, nil
}
