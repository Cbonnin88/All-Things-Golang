package main

import (
	"fmt"
	"sort"
)

func main() { //nolint:typecheck

	slice1 := []float64{2.3,10.5,8.7,3.2,4.7,1.9,5.5,7.9,6.3,9.6}
	slice2 := []string{"House","Truck","Couch","Song","Bird","Fart"}

	fmt.Println("Unsorted:",slice1)
	sort.Float64s(slice1)
	fmt.Println("Sorted:",slice1)
	fmt.Println("")

	fmt.Println("Unsorted",slice2)
	sort.Strings(slice2)
	fmt.Println("Sorted:",slice2)

}
