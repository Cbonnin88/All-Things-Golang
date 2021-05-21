package main

import "fmt"

/* The struct here is called 'car'
type car struct {
	model string
	year int
	transmission string
	new bool
} */

func main() { //nolint:typecheck
	// This is an anonymous struct because it doesn't have a name, we are just created a struct with its fields:
	car1 := struct {
		model        string
		year         int
		transmission string
		new          bool
	}{
		model:        "Ford",
		year:         2000,
		transmission: "Automatic",
		new:          true,
	}
	fmt.Println(car1)

}
