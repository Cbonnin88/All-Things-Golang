package main

import "fmt"

func main(){
	country := "France"
	countryAddress := &country

	fmt.Println("The address of country is: ", countryAddress)
}
