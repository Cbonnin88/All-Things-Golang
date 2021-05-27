package main

import (
	"fmt"
	"math"
)

// declaring our 'shape' interface:
type shape interface{
	area() float64
	parameter() float64

}

// declaring our structs:
type rectangle struct {
	width 		float64
	height 		float64
	isAShape	bool
}

type circle struct {
	radius float64
	isAShape bool
}

func main() { //nolint:typecheck
	r := rectangle{width: 5,height: 8,isAShape: true}
	c := circle{radius: 7,isAShape: true}


	isAShape(r,c)
	Calculate(r)
	Calculate(c)
}

// Implementing our 'area' method with a return statement
func (r rectangle) area() float64 {
	return r.height * r.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// implementing the 'parameter' method also with a return statement
func (r rectangle) parameter() float64 {
	return 2 * r.width + 2 * r.height
}

func (c circle) parameter() float64 {
	return 2 * math.Pi * c.radius
}

// our "isAShape" function
func isAShape (r rectangle, c circle) {
	fmt.Println("Is rectangle a shape ?:",r.isAShape)
	fmt.Println("Is circle a shape ?:",c.isAShape)
}


// Calculate function passes in our 'shape' interface:
func Calculate (s shape){
	fmt.Println(s.area())
	fmt.Println(s.parameter())
}
