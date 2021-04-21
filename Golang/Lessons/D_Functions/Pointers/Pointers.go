package main

import "fmt"

func main() {
	country := "France"
	countryAddress := &country
	fmt.Println("The address of country is: ", country, countryAddress)

	pointer := 7
	a := &pointer
	fmt.Println("The address of number seven is: ", pointer, a)

	pointer2 := "Go"
	b := &pointer2
	fmt.Println("The address of our string Go is : ", pointer2, b)
}

//
