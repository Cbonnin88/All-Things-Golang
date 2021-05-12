package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter a number: ")

	var num int
	_, err := fmt.Scan(&num)
	if err != nil {
		return
	}
	fmt.Printf("We have assigned the variable %v the type %T\n\n", num, num)
	fmt.Printf("Decimal form: %d\n", num)
	fmt.Printf("Binary form: %b\n", num)
	fmt.Printf("Hex form: %#x\n\n", num)

	shift := num << 1
	fmt.Println("Bit shifting our int:")
	fmt.Printf("Decimal form (bit Shifting): %d\n", shift)
	fmt.Printf("Binary form (bit shifting): %b\n", shift)
	fmt.Printf("hex fomr (bit shifting): %#x\n", shift)

}
