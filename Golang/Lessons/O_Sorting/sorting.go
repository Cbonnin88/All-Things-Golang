package main

import (
	"fmt"
	"sort"
)


func main() {
	// Sorting a slice of int:
	slice := []int{2,5,3,4,1,6,7}
	fmt.Println("Our original slice:",slice)
	sort.Ints(slice)
	fmt.Println("Our sorted Slice:",slice)
	fmt.Println("")

	// Sorting a slice of String:
	slice2 := []string{"Karl", "Alcina","Ethan","Harnkess", "Donna","Miranda","Agatha","Mia","Bolognaise"}
	fmt.Println("our original slice:",slice2)
	sort.Strings(slice2)
	fmt.Println("our sorted slice:",slice2)
	fmt.Println("")

	slice3 := []float64{2.3,1.1,10.8,6.7,5.5,9.8,3.4,4.2,8.3,7.9}
	fmt.Println("our original slice:",slice3)
	sort.Float64s(slice3)
	fmt.Println("Our sorted slice:",slice3)

}
