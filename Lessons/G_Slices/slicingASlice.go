package main

import "fmt"

func main() {
	slice1 := []string{"France", "Germany", "Spain", "Finland", "Austria"}
	// We can slice a slice using the "colon (:) operator:
	fmt.Println("Our entire slice:", slice1[0:])
	fmt.Println("Our slice starting from index 1", slice1[1:])
	fmt.Println("Our index 2 and 4:", slice1[2:4])
	fmt.Println("Our index 0 and 3:", slice1[0:3])
	fmt.Println("")

	// Looping without a "range" clause:
	fmt.Println("Looping without using a range clause:")
	slice2 := []int{5, 10, 15, 20, 25, 30}
	for i := 0; i <= 5; i++ {
		fmt.Println(i, slice2[i])
	}

}
