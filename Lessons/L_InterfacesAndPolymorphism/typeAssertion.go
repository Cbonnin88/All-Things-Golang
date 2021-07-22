package main

import (
	"fmt"
	"math"
)

// declaring our 'shape' interface:
type shape2 interface{
	area2() float64
	parameter2() float64
}

type polygon interface {
	Getsides() int
}

// declaring our structs:
type line struct {
	radius float64
	isAShape bool
	sides	 int
}

func main() { //nolint:typecheck
	 l := line{0, false, 0}
	 Calculate2(l)
}
func (l line) area2() float64 {
	fmt.Println("The area of a line:")
	return math.Pi * l.radius * l.radius
}

func (l line) parameter2() float64 {
	fmt.Println("The parameter of a line:")
	return 2 * math.Pi * l.radius
}

func (l line) Getsides() int {
	fmt.Println("The number of sides a rectangle has:")
	return l.sides
}

func Calculate2 (s shape2) {
	// Type assertion
	 rect := s.(polygon)
	 fmt.Printf("My type is: %T",rect)

}

/*
**Assertions are for interface types**

We can also use type assertion via an 'if' statement:
if rect, ok := s.(line); ok {
	fmt.Println(rect.parameter2())
	fmt.Printf("My type is %T", rect)
}
*/
