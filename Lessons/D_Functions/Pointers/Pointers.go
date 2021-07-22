package main

import "fmt"

func main() {
	country := "France"
	fmt.Println("The address of country is:", country, &country)
	fmt.Printf("our 'country' variable is of type: %T\n",country)
	fmt.Printf("our pointer is of type: %T\n",&country)
	fmt.Println("")

	pointer := 7
	fmt.Println("The address of number seven is:", pointer, &pointer)
	fmt.Printf("Our variable is of type: %T\n",pointer)
	fmt.Printf("Our pointer is of type: %T\n",&pointer)
	fmt.Println("")

	pointer2 := "Go"
	fmt.Println("The address of our string Go is :", pointer2, &pointer2)
	fmt.Printf("Our variable is of type: %T\n",pointer2)
	fmt.Printf("Our pointer is of type: %T\n",&pointer2)
	fmt.Println("")

	pointer3 := 123.454
	fmt.Println("The address of our float is:",pointer3,&pointer3)
	fmt.Printf("Our variable is of type: %T\n",pointer3)
	fmt.Printf("Our pointer is of type: %T\n",&pointer3)


}

// A pointer is a memory address
// '&' gives you the address
// '*' gives you the value stored at an address when you have the address
// Once again, EVERYTHING IN GO IS PASSED BY VALUE

