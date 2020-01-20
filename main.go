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

func sortBy(houses []house, less func(a, b house) bool) []house {
	result := make([]house, len(houses))
	copy(result, houses)
	sort.Slice(result, func(i, j int) bool {
		return less(result[i], result[j])
	})
	return result
}
func filterBy(houses []house, compare func(a house) bool) []house {
	result := make([]house, 0)
	for _, house := range houses {
		if compare(house) {
			result = append(result, house)
		}
	}
	return result
}

func sortByPriceDesc(houses []house) []house {
	return sortBy(houses, func(a, b house) bool {
		return a.price > b.price
	})
}

func sortByPriceAsc(houses []house) []house {
	return sortBy(houses, func(a, b house) bool {
		return a.price < b.price
	})
}




func sortByDistanceFromCentreDesc(houses []house) []house {
	return sortBy(houses, func(a, b house) bool {
		return a.distanceFromCentre > b.distanceFromCentre
	})
}

func sortByDistanceFromCentreAsc(houses []house) []house {
	return sortBy(houses, func(a, b house) bool {
		return a.distanceFromCentre < b.distanceFromCentre
	})

}



func searchByPriceLimit(houses []house, limitStart, limitEnd int64) []house {
	return filterBy(houses, func(a house) bool {
		return a.price >= limitStart && a.price <= limitEnd},
	) //
}

func searchByRegion(houses []house, region string) []house {
	return filterBy(houses, func(a house) bool {
		return a.region == region
	})
}



func searchByMaxPrice(houses []house, limit int64) []house {
	return filterBy(houses, func(h house) bool {
		return h.price <= limit
	})
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

