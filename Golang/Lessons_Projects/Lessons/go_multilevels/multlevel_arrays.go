package main

import "fmt"

func main() {
	// Multi-Dimensional arrays, are the arrays of arrays of the same type.
	// In Go language, you can create a multi-dimensional array using the following syntax:
	// Array_name =[Length1][Length2]type{………},

	var countries = [4][4]string{{"France", "Germany", "Denmark", "Finland"},
		{"Canada", "United States", "Mexico", "Trindad and Tabago"},
		{"Venezuela", "Uruguay", "Colombia", "Chile"},
		{"New Zealand", "Australia", "Japan", "South Korea"}}

	fmt.Println("Elements of Array: ")
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			fmt.Println(countries[x][y])
		}
	}

}
