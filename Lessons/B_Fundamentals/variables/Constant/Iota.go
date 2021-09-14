package main

import "fmt"

func main() {
	fmt.Println("Using iota:")
	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println("Iota can automatically increment:")
	const (
		e = iota
		f
		g
		h
	)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)

	fmt.Println("Example of reiterating iota:")
	const (
		num1 = iota
		num2
		num3
	)
	const (
		num4 = iota
		num5
		num6
	)
	fmt.Println(num1, "\n", num2, "\n", num3, "\n", num4, "\n", num5, "\n", num6)
	fmt.Println("incrementation of iota:")
	const (
		num7 = iota + 1
		num8 = iota + 2
		num9 = iota + 3
	)
	fmt.Println(num7, ": increments by one")
	fmt.Println(num8, ": Increments by two")
	fmt.Println(num9, ": Increments by three")
}
