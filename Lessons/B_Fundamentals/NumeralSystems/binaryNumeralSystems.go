package main

import "fmt"

func main() {
	s := "C"
	fmt.Println(s)
	fmt.Println("Converting to a slice of byte:")
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Println("Accessing our index position:")
	a := bs[0]
	fmt.Println(a)
	fmt.Printf("this is a %T type\n", a)
	fmt.Printf("In Binary: %b\n", a)
	fmt.Printf("In hexadecimal: %#x\n", a)

}
