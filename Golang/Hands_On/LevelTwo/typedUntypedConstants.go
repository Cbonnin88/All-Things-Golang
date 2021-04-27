package main

import "fmt"

func main() { //nolint:typecheck
	const (
		word1 string  = "I am a typed constant"
		num1  int     = 21
		dec1  float64 = 55.77
	)

	fmt.Println(word1, "- This is a  typed string constant")
	fmt.Println(num1, "- This is a typed integer constant")
	fmt.Println(dec1, "- This is a typed float constant\n")

	const a = "house"
	const b = 23
	const c = false
	fmt.Println(a, "- This is an untyped string constant")
	fmt.Println(b, "- This is an untyped integer constant")
	fmt.Println(c, "- This is an untyped boolean constant")

}
