package main

import "fmt"

func main() {
	// creating a Map
	var countryCapital map[string]string = map[string]string{
		"France":  "Paris", // "France" represent the value and "Paris" represents the key"
		"Germany": "Berlin",
		"Japan":   "Tokyo",
		"Spain":   "Madrid",
	}
	fmt.Println("Creating a Map:\n ", countryCapital)
	// Another way to create a map
	countryCapital2 := make(map[string]string)
	fmt.Println("Creating an empty Map:\n", countryCapital2) // creates and empty map

	// accessing a map value
	fmt.Println("Accessing a map value: ", countryCapital["France"], ",", countryCapital["Germany"])

	// changing a map value
	countryCapital["Japan"] = "Kyoto"
	fmt.Println("Changing a map value: ", countryCapital)

	// Adding a value
	countryCapital2["United States"] = "Washington DC"
	fmt.Println("Adding a value to countryCapital2: ", countryCapital2)

	// Deleting a value
	delete(countryCapital2, "United States")
	fmt.Println("Deleting a value from countryCapital2: ", countryCapital2)

	// checking if a value exists
	val, ok := countryCapital["Germany"]
	fmt.Println("Does the value 'Germany' exist ?: ", val, ok)

	// using the len() method for maps
	fmt.Println("This map has", len(countryCapital), "keys inside of it")
}
