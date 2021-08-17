package main

import (
	"fmt"
	"strconv"
)

func main() {
	strToIntConversion()
	intToStringConversion()

}

func strToIntConversion(){
	fmt.Println("String to Integer Conversion:")

	var num int
	// Here we will use the strconv package to convert our string:
	num, err := strconv.Atoi("12345")


	// Handling our error:
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(num + 1) // will give us 1234568
}

func intToStringConversion() {
	fmt.Println("Integer to String:")


	word := strconv.Itoa(1234567)
	fmt.Println(word)
}
