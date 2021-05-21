package main

import (
	"fmt"
)

// a SLICE allows you to group together VALUES of the same TYPE

func main() { //nolint:typecheck
	// Slicing an Array/Slice
	arr := [5]int{101, 202, 316, 412, 100} // an Array that we will slice
	slice1 := arr[:]                       // will print entire array
	slice2 := arr[1:2]                     // will print only index 1
	slice3 := arr[0:2]                     // will print indexes 0 and 1
	slice4 := arr[1:4]                     // will print indexes 1, 2 and 3
	cap1 := arr[0:4]                       // will print indexes 0, 1, 2, 3 and 4
	fmt.Println("This prints the entire array/slice: ", slice1)
	fmt.Println("Index 1: ", slice2)
	fmt.Println("Indexes 0 and 1: ", slice3)
	fmt.Println("Indexes 1, 2 and 3: ", slice4)
	fmt.Println("The capacity of this slice is: ", cap(cap1)) // cap() shows the slices capacity, the capacity for this slice is 5
	fmt.Println("Our extended slice: ", slice4[:cap(slice4)]) // this will extend the slices so that is is the full length of the array.

	fmt.Println("")

	fmt.Print("Making a Slice:\n")
	// Making a Slice:
	var sliceEx = []int{6, 12, 18, 24, 30, 36} // this is a slice of int
	fmt.Println("Our slice: ", sliceEx, "\nThe capacity of this slice is: ", cap(sliceEx))
	fmt.Println("This will give us the same capacity: ", cap(sliceEx[:4]), "\n")
	// we will get 6 again because the slice is still pointing to our underlying array
}
