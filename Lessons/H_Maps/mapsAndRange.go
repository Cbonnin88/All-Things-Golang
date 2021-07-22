package main

import "fmt"

func main() { //nolint:typecheck

	// Ranging a map

	// Creating our map
	var employee = map[string]int{
		"Anna" : 45,
		"Mark" : 30,
		"Sylvia": 60,
		"Tom" : 18,
	}

	// looping over and ranging our map:
	fmt.Println("Looping and using 'range' with our map:")
	for  k, v := range employee {
		fmt.Println(k,v)

	}

}
