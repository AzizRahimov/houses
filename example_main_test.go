package main

import "fmt"

var houses = [] house{
	{
		id:                 1,
		name:               "2 комнатная квартира",
		price:              100_000,
		region:             "Фирдавси",
		distanceFromCentre: 10,
	},
	{
		id:                 2,
		name:               "3 комнатная квартира",
		price:              140_000,
		region:             "Сино",
		distanceFromCentre: 8,
	},
	{
		id:                 3,
		name:               "5 комнатная квартира",
		price:              500_000,
		region:             "Шохмансур",
		distanceFromCentre: 14,
	},
	{
		id:                 4,
		name:               "2 комнатная квартира",
		price:              200_000,
		region:             "Шохмансур",
		distanceFromCentre: 15,
	},
	
}

var housesEmpty = []house{}

func ExampleSortByPriceDesc()  {
	result :=sortByPriceDesc(houses)
	fmt.Println(result)
	// Output: [{3 5 комнатная квартира 500000 Шохмансур 14} {4 2 комнатная квартира 200000 Шохмансур 15} {2 3 комнатная квартира 140000 Сино 8} {1 2 комнатная квартира 100000 Фирдавси 10}]

}
func ExampleSortByPriceAsc()  {
	result := sortByPriceAsc(houses)
	fmt.Println(result)
	// Output: [{1 2 комнатная квартира 100000 Фирдавси 10} {2 3 комнатная квартира 140000 Сино 8} {4 2 комнатная квартира 200000 Шохмансур 15} {3 5 комнатная квартира 500000 Шохмансур 14}]
}
func ExampleSortByDistanceFromCentreDesc()  {
	result := sortByDistanceFromCentreDesc(houses)
	fmt.Println(result)
	//Output:[{4 2 комнатная квартира 200000 Шохмансур 15} {3 5 комнатная квартира 500000 Шохмансур 14} {1 2 комнатная квартира 100000 Фирдавси 10} {2 3 комнатная квартира 140000 Сино 8}]

}
func ExampleSortByDistanceFromCentreAsc()  {
	result := sortByDistanceFromCentreAsc(houses)
	fmt.Println(result)
	//Output: [{2 3 комнатная квартира 140000 Сино 8} {1 2 комнатная квартира 100000 Фирдавси 10} {3 5 комнатная квартира 500000 Шохмансур 14} {4 2 комнатная квартира 200000 Шохмансур 15}]
}

func ExampleSearchByMaxPrice_ManyResults()  {
	result := searchByMaxPrice(houses,200_000)
	fmt.Println(result)
	// Output: [{1 2 комнатная квартира 100000 Фирдавси 10} {2 3 комнатная квартира 140000 Сино 8} {4 2 комнатная квартира 200000 Шохмансур 15}]
}
func ExampleSearchByMaxPrice_NoResult()  {
	result := searchByMaxPrice(housesEmpty,100000000)
	fmt.Println(result)
	// Output: []
}
func ExampleSearchByPriceLimit_ManyResults()  {
	result := searchByPriceLimit (houses, 100_000, 300_000 )
	fmt.Println(result)
	//Output: [{1 2 комнатная квартира 100000 Фирдавси 10} {2 3 комнатная квартира 140000 Сино 8} {4 2 комнатная квартира 200000 Шохмансур 15}]
}

func ExampleSearchByPriceLimit_NoResults()  {
	result := searchByPriceLimit (housesEmpty, 1000_000, 3000_000 )
	fmt.Println(result)
	//Output: []
}


func ExampleSearchByRegionsManyResults()  {
	result := searchByRegions (houses, []string{"Фирдавси", "Сино"})
	fmt.Println(result)
	//Output: [{1 2 комнатная квартира 100000 Фирдавси 10} {2 3 комнатная квартира 140000 Сино 8}]
}

func ExampleSearchByRegions_NoResults()  {
	result := searchByRegions (housesEmpty, []string{"Фирдавси", "Сино"})
	fmt.Println(result)
	//Output: []
}

func ExampleSearchByRegion_ManyResults()  {
result := searchByRegion(houses, "Сино")
fmt.Println(result)
//Output: [{2 3 комнатная квартира 140000 Сино 8}]
}
func ExampleSearchByRegion_NoResults()  {
	result := searchByRegion(housesEmpty, "Сино")
	fmt.Println(result)
	//Output: []
}


