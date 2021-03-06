package main

import "fmt"

// Printf can interpolate stings, or leave placeholders in a string and use values to fill in the placeholder
/*
Example:
guess:= "C"
fmt.Printf("Is %v your final answer?)
Prints : "Is C your final answer?
*/

func main() {
	guess := "Chartres" // String value = %v
	age := 12           // Integer value = %d
	gpa := 4.5          // float or double value = %f
	euro := 12.65       // float or double value =%f (for decimal places %.2f)
	fmt.Printf("Is that the city of %v up ahead?\n", guess)
	fmt.Printf("My father is turning %d tomorrow\n", age)
	fmt.Printf("You average is about: %f\n", gpa)
	fmt.Printf("In France, I only paid €%.2f for a arm caste!!!!\n", euro)

}

// As long as you have enough arguments, you can even add multiple placeholders
/*
Example:
selection := "soup"
selection := "salad"
fmt.Println("Do you want %v or %v tonight?")
Prints: Do you want soup or salade tonight
*/
