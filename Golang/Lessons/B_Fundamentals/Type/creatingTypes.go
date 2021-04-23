package main

import "fmt"

func main() {

	// With Go, it's possible to create your own Types:
	// ** Side note: %T shows us the 'type' that is our value **

	// Example of a int type using %T:
	var num int
	num = 32
	fmt.Printf("32 is an %T type\n", num)

	// Creating our own type:
	type batman float32 // this tell us that our type is 'batman' and his underlying type is a float

	// declaring our variable the type 'batman' :
	var b batman = 25.78
	fmt.Println("Declaring our 'b' variable of type 'batman':", b)
	fmt.Printf("%T: this says that 25.78 is of type 'batman' from the main\n", b) // 25.78 is of type 'batman'

	// Example 2:
	type catwoman string
	var c catwoman = "meow"
	fmt.Println(c)
	fmt.Printf("%T", c)

}
