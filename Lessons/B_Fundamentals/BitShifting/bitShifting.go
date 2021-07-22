package main

import "fmt"

func main() {
	// bit shifting is when you shift binary digits to the left or right
	fmt.Println("Decimal and binary before bit shifting:")
	num1 := 12
	fmt.Printf("Decimal: %d\t\t Binary: %b\n", num1, num1)
	fmt.Println("Bit Shifting example:")
	bitShift := num1 << 1
	fmt.Printf("Decimal(with Shifting): %d\t\t Binary(with Shifting): %b\n\n", bitShift, bitShift)

	fmt.Println("Kilobytes, Megabytes, Gigabytes:")
	kb := 1024
	mb := kb * 1024
	gb := mb * 1024
	fmt.Printf("Kilobyte = Decimal: %d\t\t\tBinary: %b\n", kb, kb)
	fmt.Printf("Megabyte = Decimal: %d\t\t\tBinary: %b\n", mb, mb)
	fmt.Printf("Gigabyte = Decimal: %d\t\t\tBinary: %b\n\n", gb, gb)

	fmt.Println("Shifting with iotas:")
	const (
		_   = iota
		kb2 = 1 << (iota * 10) // shifting over ten times
		mb2 = 1 << (iota * 10) // shifting over 20 times
		gb2 = 1 << (iota * 10) // shifting over 30 times
	)

	fmt.Printf("Kilobyte = Decimal: %d\t\t\tBinary: %b\n", kb2, kb2)
	fmt.Printf("Megabyte = Decimal: %d\t\t\tBinary: %b\n", mb2, mb2)
	fmt.Printf("Gigabyte = Decimal: %d\t\t\tBinary: %b\n\n", gb2, gb2)
}
