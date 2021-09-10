package main

import "fmt"

func main() {

	// Deleting from a slice:

	slice1 := []string{"Paris", "Berlin", "Washington DC", "Perville", "Sockoland"}
	slice2 := []string{"Madrid", "Troutston", "Tokyo", "FreetheLandstan", "Turkeyburg"}

	slice1 = append(slice1, slice2...)
	fmt.Println("Our slices combined with real and fake cities:\n", slice1)
	fmt.Println("")

	// to delete an element in a slice, we use the append() method:
	slice1 = append(slice1[:3], slice2[0], slice2[2])
	fmt.Println("Our slice with only the real cities:\n", slice1)
	fmt.Println("")

	slice3 := []string{"Helsinki", "Panam", "Pandamonium", "Tel Aviv", "Soeul"}
	slice4 := []string{"Beijing", "Porkton", "Ottawa", "Canberra", "NewBeeftonville"}

	slice3 = append(slice3[1:3], slice4[1], slice4[4])
	fmt.Println("Our slice with fake cities:\n", slice3)

}
