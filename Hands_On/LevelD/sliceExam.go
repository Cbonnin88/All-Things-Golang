package main

import "fmt"

func main() {
	var slice = []int{5, 10, 15, 20, 25, 30, 35, 40, 45, 50}

	// Ranging over our slice:
	for i, v := range slice {
		fmt.Println(i, v)
	}
	fmt.Printf("Our slice type: %T\n", slice)
	fmt.Println("")

	slice2 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(slice2[:5])
	fmt.Println(slice2[5:10])
	fmt.Println(slice2[2:7])
	fmt.Println(slice2[1:6])

}
