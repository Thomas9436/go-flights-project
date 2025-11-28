package sorters

import (
	"aggregator/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func mk(dep, arr string, price float64) model.Flight {
	d, _ := time.Parse(time.RFC3339, dep)
	a, _ := time.Parse(time.RFC3339, arr)
	return model.Flight{DepartureTime: d, ArrivalTime: a, Price: price}
}

func TestSortByPriceAsc(t *testing.T) {
	f := []model.Flight{
		mk("2026-01-01T10:00:00Z", "2026-01-01T12:00:00Z", 800.0),
		mk("2026-01-02T10:00:00Z", "2026-01-02T12:00:00Z", 400.0),
		mk("2026-01-03T10:00:00Z", "2026-01-03T12:00:00Z", 1200.0),
	}
	SortByPriceAsc(f)
	assert.Equal(t, 400.0, f[0].Price)
	assert.Equal(t, 800.0, f[1].Price)
	assert.Equal(t, 1200.0, f[2].Price)
}

func TestSortByDepartureAsc(t *testing.T) {
	f := []model.Flight{
		mk("2026-01-03T10:00:00Z", "2026-01-03T12:00:00Z", 1),
		mk("2026-01-01T10:00:00Z", "2026-01-01T12:00:00Z", 1),
		mk("2026-01-02T10:00:00Z", "2026-01-02T12:00:00Z", 1),
	}
	SortByDepartureAsc(f)
	assert.True(t, f[0].DepartureTime.Before(f[1].DepartureTime))
	assert.True(t, f[1].DepartureTime.Before(f[2].DepartureTime))
}

func TestSortByTravelTimeAsc(t *testing.T) {
	f := []model.Flight{
		mk("2026-01-01T10:00:00Z", "2026-01-01T20:00:00Z", 1), // 10h
		mk("2026-01-01T10:00:00Z", "2026-01-01T13:00:00Z", 1), // 3h
		mk("2026-01-01T10:00:00Z", "2026-01-01T15:00:00Z", 1), // 5h
	}
	SortByTravelTimeAsc(f)
	assert.Equal(t, 3*time.Hour, f[0].ArrivalTime.Sub(f[0].DepartureTime))
	assert.Equal(t, 5*time.Hour, f[1].ArrivalTime.Sub(f[1].DepartureTime))
}
