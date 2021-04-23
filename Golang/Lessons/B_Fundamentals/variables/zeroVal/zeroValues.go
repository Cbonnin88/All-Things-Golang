package main

import "fmt"

func main() {
	var emptyClassInt int       // Prints 0
	var emptyGrace float32      // Prints 0
	var emptyTeacherName string // Doesn't Print anything
	var isPassFail bool         // Prints false

	// All numeric variables have a value of "0" before assignment
	// String variables have a default value of ""

	/* var emptyValue int
	   var emptyValue int = 0
	   These two are the same */

	fmt.Println(emptyClassInt, emptyGrace, emptyTeacherName, isPassFail)

}
