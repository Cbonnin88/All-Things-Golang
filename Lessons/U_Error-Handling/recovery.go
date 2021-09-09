package main

import "fmt"

/*
 Recover is a built-in function that regains control of a panicking goroutine.
 ** Recover is only useful inside deferred functions **
*/

func main() {
	citySlice([]string{"Paris", "Philadelphia", "Chessy", "Chartres", "Berlin", "Helsinki", "Oslo", "London", "Windsor", "Nantes", "Rennes"}, 0)
}

func citySlice(slice []string, index int) {
	defer func() {
		if p := recover(); p != nil { // here we call recover
			fmt.Printf("internal error: %v ", p) // And here we print the recover value
		}
	}()

	fmt.Printf("item %d, value %v \n", index, slice[index])
	defer fmt.Printf("defer %d \n", index)
	citySlice(slice, index+1)
}
