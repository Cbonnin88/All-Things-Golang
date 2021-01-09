package main

import "fmt"

func main(){
	star := "Polaris"
	starAddress := &star

	*starAddress = "Sirius"

	fmt.Println("We have changed the valule of star from Polaris to",star, "using dereferencing.")
}
