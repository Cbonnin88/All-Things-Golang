package main

import "fmt"

func main() {

	// make is a built-in function

	// Creating slices with the keyword "make()"
	fmt.Println("Creating a Slice using the 'make()' function:")
	sliceEx := make([]string, 6) // will give us an empty Slice
	fmt.Println("An empty Slice: ", sliceEx)
	fmt.Printf("This is our slice type: %T", sliceEx) // shows use the data type for our slice

	fmt.Println("\n")

	// A slice with len() and cap() functions:
	sliceEx2 := make([]int, 5, 9)
	fmt.Println("our slice using 'make()' again:", sliceEx2)
	fmt.Println("slice length: ", len(sliceEx2))
	fmt.Println("slice capacity:", cap(sliceEx2))

	// filling in our slice:
	sliceEx2[2] = 36
	sliceEx2[4] = 56
	fmt.Println("Filling in our slice:", sliceEx2)

	fmt.Println("")

	// appending the length of our slice:
	fmt.Println("our original slice length:", len(sliceEx))
	sliceEx = append(sliceEx, "house", "cat", "fat")
	fmt.Println("our slice using 'append()': ", len(sliceEx))

}
