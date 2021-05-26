package main

import "fmt"

func main() { //nolint:typecheck
	// A numeric type represents sets of integer or floating point values
	// With int we have unassigned and assigned values, signed values can be negative, unassigned can only be positive:

	// Integers:
	var num1 int = 77
	var num2 uint = 34
	fmt.Println("Integers:")
	fmt.Println("our unassigned int is", num1)
	fmt.Printf("this is a %T type\n", num1)
	fmt.Println("our assigned int is", num2)
	fmt.Printf("This is a %T type\n\n", num2)

	// Floats:
	fmt.Println("Floats:")
	var decimal float32 = 23.44
	var decimal2 float64 = 234.76
	fmt.Println("the value of our float is", decimal)
	fmt.Printf("'decimal' is a %T type", decimal)
	fmt.Println("the value of our second float is", decimal2)
	fmt.Printf("'decimal2' is a %T type", decimal2)
	// byte is an alias for uint8
	// rune is an alias for int32

}
