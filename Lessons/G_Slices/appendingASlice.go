package main

import "fmt"

func main() { //nolint:typecheck

	// We can add an element inside of a slice using append()

	// Using append():
	fmt.Print("Adding Elements to a Slice:\n")
	sliceEx := []string{"House", "Cat", "Kid", "Dog"}
	addOn := append(sliceEx, "Car") // append() will return a new slice with Car added to it
	fmt.Println("Our original slice: ", sliceEx)
	fmt.Println("our new slice using the append() method: ", addOn)

	fmt.Println("")

	// if you want to append the slice using the same variable name:
	fmt.Println("Appending a slice without create a separate variable:")
	sliceEx2 := []string{"ghost", "feet", "house", "clean"}
	fmt.Println("our original slice: ", sliceEx2)
	sliceEx2 = append(sliceEx2, "game")
	fmt.Println("our new slice: ", sliceEx2)

	fmt.Println("")

	// Adding on or appending a slice:
	fmt.Println("Appending two slices:")
	slice1 := []int{100, 200, 300, 400, 500}
	slice2 := []int{10, 20, 30, 40, 50}
	slice1 = append(slice1, slice2...)
	fmt.Println(slice1)

}
