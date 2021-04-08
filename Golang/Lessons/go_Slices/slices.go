package main

import (
	"fmt"
)

// Slices are a different data type than Arrays

func main() {
	// Slicing an Array:
	var arr [5]int = [5]int{101, 202, 316, 412, 100} // an Array that we will slice
	var slice1 []int = arr[:]                        // will print entire array
	var slice2 []int = arr[1:2]                      // will print only index 1
	var slice3 []int = arr[0:2]                      // will print indexes 0 and 1
	var slice4 []int = arr[1:4]                      // will print indexes 1, 2 and 3
	var cap1 []int = arr[0:4]                        // will print indexes 0, 1, 2, 3 and 4
	fmt.Println("This prints the entire array: ", slice1)
	fmt.Println("Index 1: ", slice2)
	fmt.Println("Indexes 0 and 1: ", slice3)
	fmt.Println("Indexes 1, 2 and 3: ", slice4)
	fmt.Println("The capacity of this slice is: ", cap(cap1))       // cap() shows the slices capacity, the capacity for this slice is 5
	fmt.Println("Our extended slice: ", slice4[:cap(slice4)], "\n") // this will extend the slices so that is is the full length of the array.

	fmt.Print("Making a Slice:\n")

	// Making a Slice:
	var sliceEx []int = []int{6, 12, 18, 24, 30, 36} // this is a slice of int
	fmt.Println("Our slice: ", sliceEx, "\nThe capacity of this slice is: ", cap(sliceEx))
	fmt.Println("This will give us the same capacity: ", cap(sliceEx[:4]), "\n")
	// we will get 6 again because the slice is still pointing to our underlying array

	fmt.Print("Adding Elements to a Slice:\n")
	var sliceEx2 []string = []string{"House", "Cat", "Kid", "Dog"}
	addOn := append(sliceEx2, "Car") // append() will return a new slice with Car added to it
	fmt.Println("Our original slice: ", sliceEx2)
	fmt.Println("our new slice using the append() method: ", addOn, "\n")

	// if you want to append the slice using the same variable name:
	fmt.Println("Appending a slice without create a separate variable:")
	var sliceEx3 []string = []string{"ghost", "feet", "house", "clean"}
	sliceEx3 = append(sliceEx3, "game")
	fmt.Println("our original slice: ", sliceEx3, "\n")

	// Creating slices with the keyword "make()"
	fmt.Println("Creating a Slice using the 'make()' function:")
	sliceEx4 := make([]string, 6) // will give us an empty Slice
	fmt.Println("An empty Slice: ", sliceEx4)
	fmt.Printf("This is our slice type: %T", sliceEx4) // shows use the data type for our slice
}
