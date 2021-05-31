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
type rectangle2 struct {
	width 		float64
	height 		float64
	isAShape	bool
	sides		int
}

type line struct {
	radius float64
	isAShape bool
}

func main() { //nolint:typecheck
	r2 := rectangle2{width: 5,height: 10, isAShape: true, sides: 4}
	// l := line{0, false},
	Calculate2(r2)
	// Calculate2(l), this will print 'main.line'
}

func (r rectangle2) area2() float64 {
	fmt.Println("the area of the rectangle:")
	return r.height * r.width
}


func (r rectangle2) parameter2() float64 {
	fmt.Println("The parameter of the rectangle:")
	return 2 * r.width + 2 * r.height
}

func (l line) area2() float64 {
	fmt.Println("The area of a line:")
	return math.Pi * l.radius * l.radius
}

func (l line) parameter2() float64 {
	fmt.Println("The parameter of a line:")
	return 2 * math.Pi * l.radius
}

func (r rectangle2) Getsides() int {
	fmt.Println("The number of sides a rectangle has:")
	return r.sides
}

func Calculate2 (s shape2) {

	// Type assertion
	 rect := s.(polygon)
	 fmt.Printf("My type is: %T",rect)

}

/*
We can also use type assertion via an 'if' statement:
if rect, ok := s.(line); ok {
	fmt.Println(rect.parameter2())
	fmt.Printf("My type is %T", rect)
}
*/
