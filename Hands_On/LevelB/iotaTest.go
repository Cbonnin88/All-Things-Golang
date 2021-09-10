package main

import "fmt"

func main() {
	const (
		y = 2021 + iota
		y2
		y3
		y4
		y5
	)
	age := y - 1988
	age1 := y2 - 1988
	age2 := y3 - 1988
	age3 := y4 - 1988
	age4 := y5 - 1988

	fmt.Println("The current year is", y, ". I am currently", age, "years old")
	fmt.Println("Next year will be", y2, ". I will be", age1, " years old")
	fmt.Println("In two years it will be", y3, ". I will be", age2, "years old")
	fmt.Println("In three years it will be", y4, ". I will be", age3, "years old")
	fmt.Println("In four years, it will be", y5, ". I will be", age4, "years old")

}
