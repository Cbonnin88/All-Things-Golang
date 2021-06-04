package main

import "fmt"

type person struct {
	firstName	string
	lastName	string
	age 		int
	isHuman     bool
}

func main() { //nolint:typecheck

	p1 := person {
		firstName: "Ethan",
		lastName:  "Winters",
		age: 		37,
		isHuman:   true,
	}
	fmt.Println("Before:")
	fmt.Println(p1)
	fmt.Println("")

	changeMe(&p1)

	fmt.Println("After:")
	fmt.Println(p1)

}

func changeMe(p *person) {
	p.firstName = "Alcina"
	p.lastName = "Dimitrescu"
	p.age = 107
	p.isHuman = false

	(*p).firstName = "Alcina"
	(*p).lastName =  "Dimitrescu"
	(*p).age = 107
	(*p).isHuman = false
}
