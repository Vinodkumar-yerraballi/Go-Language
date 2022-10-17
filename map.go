package main

import (
	"fmt"
)

func main() {
	CountryPopulation := map[string]int{
		"India":   123,
		"USA":     445,
		"Germany": 786,
		"China":   9999,
	}
	fmt.Println(CountryPopulation)

	country := make(map[string]int, 5)
	fmt.Println(country["apple"])
	// how to add name to the map function
	CountryPopulation["Uk"] = 5435
	fmt.Println(CountryPopulation)
	// how to delete element in map
	delete(CountryPopulation, "China")
	fmt.Println(CountryPopulation)
	fmt.Println(CountryPopulation["India"])
	// check the value is present in the map or not it give false value because it's not in the list

	temp, ok := CountryPopulation["Nepal"]
	fmt.Println(temp, ok)
	_, ok1 := CountryPopulation["USA"]
	fmt.Println(ok1)
}
