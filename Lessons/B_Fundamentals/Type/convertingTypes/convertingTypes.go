package main

import "fmt"

func main() {
	// In Go we can convert types from on type to another
	// Example :
	var num1 int = 23 // num1 has been declared as an int, it can store a value of int

	type school float32
	var study school = 98.78

	// our types before conversion:
	fmt.Println("Example one, conversion from 'school' to an int:")
	fmt.Println("num1 =", num1)
	fmt.Printf("'num1' is a %T type\n", num1)
	fmt.Println("study =", study)
	fmt.Printf("'study' is a %T type\n", study)

	// we convert our type 'school' to a type int
	num1 = int(study)

	// 'num1' takes on the values of 'school'
	fmt.Println("'num1' changes from 23 to", num1, ",taking on the values of 'school' but changing our float to an int") // num1 become an int of 99
	fmt.Printf("%T\n\n", num1)

}
