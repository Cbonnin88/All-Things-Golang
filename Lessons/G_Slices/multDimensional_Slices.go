package main

import "fmt"

func main() {
	slice1 := [][]int{
		{10, 20},
		{30, 40},
		{50, 60},
		{70, 80},
	}
	fmt.Println("Slice 1 MultiDimensional: ", slice1)

	// Creating a MD Slice using a loop:
	slice2 := [3][2]string{
		{"France, Germany"},
		{"Spain, Finland"},
		{"Belgium, The Netherlands"},
	}
	fmt.Println("MultiDimensional Slice:")
	for i := 0; i < len(slice2); i++ {
		fmt.Println(slice2[i])
	}

	// Creating an MD slice without the brackets:
	slice3 := [][]int{
		{10, 40},
		{20, 50},
		{30, 60},
	}
	length := len(slice3)
	m := len(slice3[0])

	for i := 0; i < length; i++ {
		for h := 0; h < m; h++ {
			fmt.Println(slice3[i][h])
		}
		fmt.Print("\n")
	}

	// Creating an MD Slice with two different slices
	city := []string{"London", "Paris", "Berlin", "Washington DC"}
	country := []string{"United Kingdom", "France", "Germany", "United States"}
	geo := [][]string{city, country}
	fmt.Println(geo)
}
