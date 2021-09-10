package main

import "fmt"

type number int

var a number
var b int

func main() {
	a = 120
	b = 220

	fmt.Println(a)
	fmt.Printf("'a' is a %T type\n", a)
	fmt.Println(b)
	fmt.Printf("'b' is an %T type\n", b)

	a = 1000
	fmt.Println("'a' changed from 120 to", a)

	b = int(a)
	fmt.Println("'b' has a value of", b)
	fmt.Printf("'b' is a %T type", b)

}
