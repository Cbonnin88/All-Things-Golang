package main

import "fmt"

func main() {

	// DeReferencing or indirecting is when we use our pointer to access the address and change its value

	// '&' gives you the address
	// '*' gives you the value stored at an address when you have the address

	star := "Polaris"
	starAddress := &star
	*starAddress = "Sirius" // here we change the value of 'star' from 'Polaris' to 'Sirius'
	fmt.Println("We have changed the value of star from Polaris to", star, "using dereferencing.")
	fmt.Println("")

	// Example 2:
	legalAge := 18
	newLegalAge := &legalAge
	*newLegalAge = 21
	fmt.Println("We have changed our legal age from 18 to", legalAge)
	fmt.Println("")

	// Example 3:
	toChange := "Java"
	fmt.Println("When I started my training, I was learning", toChange)
	changeValue(&toChange)
	fmt.Println("and now I am learning", toChange)
	fmt.Println("")

	// Example 4:
	toReverse := "Tomatoes"
	fmt.Println(toReverse, "are my favorite veggies")
	replace := &toReverse
	*replace = changeValue2("onions")
	fmt.Println("But I hate", toReverse)
	fmt.Println("")

	// Example 5:
	x := 23.4
	fmt.Println("x before:", &x)
	fmt.Println("x before: ", x)
	boo(&x)
	fmt.Println("x after:", &x)
	fmt.Println("x after:", x)

}

// We can also dereference using a function:
func changeValue(str *string) {
	*str = "Go"
}
func changeValue2(str string) string {
	str = "onions"
	return str
}

func boo(y *float64) {
	fmt.Println("y before:", y)
	fmt.Println("y before:", *y)
	*y = 234.567
	fmt.Println("y after:", y)
	fmt.Println("y after:", *y)

}
