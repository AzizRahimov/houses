package main

import (
	"sort"
)

type house struct {
	id                 int64
	name               string
	price              int64
	region             string
	distanceFromCentre int64
}

func sortByExpensiveHouse(houses []house) []house {
	result := make([]house, len(houses))
	copy(result, houses)
	sort.Slice(result, func(i, j int) bool {
		return result[i].price > result[j].price

	})
	return result
}
func sortByCheapHouse(houses []house) []house {
	result := make([]house, len(houses))
	copy(result, houses)
	sort.Slice(result, func(i, j int) bool {
		return result[i].price < result[j].price

	})
	return result

}
func sortByLongDistanceFromCentre(houses []house) []house {
	result := make([]house, len(houses))
	copy(result, houses)
	sort.Slice(result, func(i, j int) bool {
		return result[i].distanceFromCentre > result[j].distanceFromCentre

	})
	return result

}

func sortByNearDistanceFromCentre(houses []house) []house {
	result := make([]house, len(houses))
	copy(result, houses)
	sort.Slice(result, func(i, j int) bool {
		return result[i].distanceFromCentre < result[j].distanceFromCentre

	})
	return result
}
func searchByMaxPrice(houses []house, limit int64) []house {
	result := make([]house, 0)
	for _, house := range houses {
		if house.price <= limit {
			result = append(result, house)
		}
	}
	return result
}
func searchByPriceLimit(houses []house, limitStart, limitEnd int64) []house {
	result := make([]house, 0)
	for _, house := range houses {
		if house.price >= limitStart && house.price <= limitEnd {
			result = append(result, house)
		}
	}
	return result
}

func searchByRegion(houses []house, region string) []house {
	result := make([]house, 0)
	for _, house := range houses {
		if house.region == region {
			result = append(result, house)
		}
	}
	return result
}

func searchByRegions(houses []house, regions []string) []house {
	result := make([]house, 0)
	for _, house := range houses {
		for _, region := range regions {
			if house.region == region {
				result = append(result, house)
			}
		}
	}
	return result
}

func main() {

	}

