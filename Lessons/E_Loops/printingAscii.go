package main

import "fmt"

func main() {
	// Printing numbers 33 to 122 as Binary, Hex and Unicode:
	for i := 33; i <= 122; i++ {
		fmt.Printf("Binary: %v\tHex: %#x\tUnicode: %#U\n", i, i, i)
	}

}
