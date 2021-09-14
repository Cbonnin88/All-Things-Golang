package main

import "fmt"

func main() {

	// A string type represents a set of string values
	// A string value is a, possibly empty sequences of bytes

	// Creating a simple String:
	var word string = "Hi there I am stringy"
	fmt.Println(word)
	fmt.Printf("I am a %T type\n\n", word)

	// A raw string literal
	rawString := `This is an example of a raw string literal`
	fmt.Println(rawString, "\n")

	// Converting a string to a byte:
	fmt.Println("Converting a string into a byte:")
	bs := []byte(word)
	fmt.Println("'word' variable as a byte:", bs)
	fmt.Printf("'bs' is now a %T type\n\n", bs)

	//  A string is made up of a slice of bytes

	fmt.Println("Find the UTF-8 characters of our string: ")
	// using a 'for' loop:
	for i := 0; i < len(word); i++ {
		fmt.Printf("%#U ", word[i])
	}

	// using range :
	fmt.Println("Using the range keyword:")
	for i, r := range word {
		fmt.Println(i, r)
	}
	fmt.Println("")
	fmt.Println("using the %@x:")
	// finding the hexadecimal
	for i, v := range word {
		fmt.Printf("at index position %d we have hex %#x\n", i, v)
	}
	fmt.Println("")
	fmt.Println("Changing the value of our string:")
	word2 := "Car"
	fmt.Printf("The orginal string is %v\n", word2)
	word2 = "House" // Assigning a new value to our string:
	fmt.Printf("now our string has been assigned the value of %v", word2)

}
