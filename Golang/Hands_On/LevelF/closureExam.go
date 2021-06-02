package main

import "fmt"

func main() { //nolint:typecheck

	a := incrementBy2()
	b := decrementby3()

	fmt.Println("Increasing:")
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println("")

	fmt.Println("Decreasing:")
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())

}

func incrementBy2() func() int {
	var multiply = 1
	return func() int {
		i := multiply * 2
		multiply++
		return i
	}

}

func decrementby3() func() int {
	var subtract = 4
	return func() int {
		n := subtract * 3
		subtract--
		return n
	}
}
