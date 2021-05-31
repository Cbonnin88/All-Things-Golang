package main

import "fmt"

type employee struct {
	firstName 	string
	lastName	string
	Title		string
	grade		int
	department	string
	salary		int
	motto		string
	isWorking	bool
}

type ceo struct {
	employee
	isCEO  bool
}

type human interface {
	saying()
}

func main() { //nolint:typecheck

	emp1 := employee{
		firstName: "Christopher",
		lastName: 	"Bonnin",
		Title: 		"Go Developer",
		grade:		1,
		salary:		28000,
		motto: 		"Working 9 to 5",
		isWorking: 	true,
	}

	emp2 := employee{
		firstName: "Alexandra",
		lastName:  "Shackles",
		Title: 		"Chief Technical Officer",
		grade:		10,
		salary:     80000,
		motto: 		"I am woman here me roar",
		isWorking:	true,
	}

	ceo1 := ceo{
		employee: employee{ // our embedded struct
			firstName: "Alcina",
			lastName:  "Dimitrescu",
			Title:		"Lady",
			grade:		11,
			salary:     100000000,
			motto:		"You stupid man-thing!!",
			isWorking:   true,
		},
		isCEO: true,
	}
	/*
	An embedded struct promotes all of it's values to the outer struct, similar to inheritance
	Inter type promotion is similar to overriding in other OOP languages
	In Go there are no classes, only TYPES
	In Go, you don't instantiate, you create a VALUE of a TYPE
	*/

	fmt.Println("Calling our struct:")
	fmt.Println(emp1)
	fmt.Println(emp2)
	fmt.Println("")

	fmt.Println("Calling individual values from the fields in our struct:")
	fmt.Println(emp1.firstName)
	fmt.Println(emp2.Title)
	fmt.Println("")

	fmt.Println("Calling our saying() function:")
	emp1.saying()
	emp2.saying()
	fmt.Println("")

	fmt.Println("Calling our saying() function sing the 'human' interface:")
	foo(emp1)
	foo(emp2)
	fmt.Println("")

	fmt.Println("Using Polymorphism in Go:")
	// This is an example of polymorphism in Go
	foo(ceo1)
	foo(emp1)
	foo(emp2)
}

func (e employee) saying(){
	fmt.Println(e.firstName, "says", e.motto)
}

func (c ceo) saying() {
	fmt.Println(c.firstName, "says", c.motto)
}

func foo (h human) {
	h.saying()
}
