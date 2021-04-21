package main

import "fmt"

func main() {
	// Example of a simple String:
	var a string
	a = "This is an example of a simple string"
	fmt.Println(a)

	// Example of a simple Integer:
	var num1 = 100
	num2 := 32
	fmt.Println("This is an example of an integer:", num1)
	fmt.Println("our second integer:", num2)

	// Example of a simple float:
	var dec = 32.2
	var dec2 = 32.34
	fmt.Println("Float 1:", dec, "\nFloat 2:", dec2)

	// Example of a simple Boolean:
	booleanValue := false
	var aBoolean bool
	fmt.Println("'aBoolean' variable will always be", aBoolean, "since we have not assigned our variable")
	fmt.Println("'true' variable will give us", booleanValue)

}

/* Go has 15 different numeric types that fall into the three categories:
int, float, and complex*/

// There are 15 different ways to describe a number in Go

/* This includes :
11 integer types
2 different floating point types
2 different complex number types*/
