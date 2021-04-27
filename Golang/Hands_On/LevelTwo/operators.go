package main

import "fmt"

func main() {
	a := (32 == 32)
	b := 120 <= 100
	c := 100 >= 99
	d := 31 != 31
	e := 10.34 < 232.22
	f := 0.99 < 0.56

	fmt.Println("32 == 32 is", a)
	fmt.Println("120 <= is", b)
	fmt.Println("100 >= 99 is", c)
	fmt.Println("31 != 31 is", d)
	fmt.Println("10.34 < 232.22 is", e)
	fmt.Println("0.99 < 0.56 is", f)

}
