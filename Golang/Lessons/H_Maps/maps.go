package main

import "fmt"

func main() { //nolint:typecheck
	// creating a Map
	var countryCapital = map[string]string{
		"France":  "Paris", // "France" represents the value and "Paris" represents the key"
		"Germany": "Berlin",
		"Japan":   "Tokyo",
		"Spain":   "Madrid",
	}
	fmt.Println("Creating a Map:\n ", countryCapital)
	fmt.Println("")

	// Another way to create a map
	countryCapital2 := make(map[string]string)
	fmt.Println("Creating an empty Map:\n", countryCapital2) // creates and empty map
	fmt.Println("")

	// accessing a map value
	fmt.Println("Accessing a map value: ", countryCapital["France"], ",", countryCapital["Germany"])
	fmt.Println("")

	// changing a map value
	fmt.Println("Changing a map value:")
	fmt.Println("Our original map:",countryCapital)
	countryCapital["Japan"] = "Kyoto"
	fmt.Println("Our new map: ", countryCapital)
	fmt.Println("")

	// Adding a value
	fmt.Println("Add a value to a map:")
	fmt.Println("Our original map:",countryCapital)
	countryCapital["United States"] = "Washington DC"
	fmt.Println("Our new map: ", countryCapital)
	fmt.Println("")

	// Deleting a value
	fmt.Println("Deleting a value in a map:")
	fmt.Println("Our original map:",countryCapital)
	delete(countryCapital, "United States")
	fmt.Println("Our new map without the United States: ", countryCapital)
	fmt.Println("")

	// checking if a value exists
	val, ok := countryCapital["Germany"]
	fmt.Println("Does the value 'Germany' exist ?: ", val, ok)
	val, ok = countryCapital["Russia"]
	fmt.Println("Does the value 'Russia' exist ?: ",val, ok)
	fmt.Println("")

	// Using an "if" statement to check if a value exists
	if val, ok := countryCapital["Finland"]; ok {
		fmt.Println(val,"is a value")
	} else {
		fmt.Println("No value exists")
	}
		fmt.Println("")

	// using the len() method for maps
	fmt.Println("This map has", len(countryCapital), "keys inside of it")
}
