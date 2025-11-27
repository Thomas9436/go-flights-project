package sorters

import (
	"aggregator/model"
	"sort"
)

// SortByPriceAsc sorts flights by price (ascending)
func SortByPriceAsc(flights []model.Flight) {
    sort.Slice(flights, func(i, j int) bool {
        return flights[i].Price < flights[j].Price
    })
}

// SortByDepartureAsc sorts by departure time (ascending)
func SortByDepartureAsc(flights []model.Flight) {
    sort.Slice(flights, func(i, j int) bool {
        return flights[i].DepartureTime.Before(flights[j].DepartureTime)
    })
}

// SortByTravelTimeAsc sorts by travel duration (arrival-departure) ascending
func SortByTravelTimeAsc(flights []model.Flight) {
    sort.Slice(flights, func(i, j int) bool {
        di := flights[i].ArrivalTime.Sub(flights[i].DepartureTime)
        dj := flights[j].ArrivalTime.Sub(flights[j].DepartureTime)
        return di < dj
    })
}

// SortByKey picks a sorter by key
func SortByKey(flights []model.Flight, key string) {
    switch key {
    case "price":
        SortByPriceAsc(flights)
    case "departure":
        SortByDepartureAsc(flights)
    case "time_travel", "travel_time":
        SortByTravelTimeAsc(flights)
    default:
        SortByPriceAsc(flights)
    }
}