package main

import (
	"fmt"
)

func main() {

	house := struct {
		location 	string
		year		int
		isAvailable	bool
		rooms		int
		floors		int
		price		int
		tva			float32
	}{
		location:	"Chartres",
		year:		2011,
		isAvailable: true,
		rooms:			6,
		floors:			3,
		price:			220000,
		tva:			20.0,
	}

	employee := struct {
		firstname	string
		lastname	string
		age 		int
		dob			string
		department	string
		grade		int
		isHired		bool
		salary		int
		taxbracket	float32
	}{
		firstname:   "Christopher",
		lastname:    "Bonnin",
		age:		  32,
		dob:		 "10/18/1988",
		department:	 "Tech/IT",
		grade:		  5,
		isHired: 	  true,
		salary: 	  28000,
		taxbracket:   4.5,
	}

	fmt.Println("our anonymous struct:", house)
	fmt.Println("our second anonymous struct",employee)
	fmt.Println("")
	fmt.Println("Calling information from our house struct:")
	fmt.Println("The location of our new house is:",house.location)
	fmt.Println("The price of our new house is:",house.price)
	fmt.Println("The year our house was built:",house.year)
	fmt.Println("The number of rooms and floors:\n","rooms:",house.rooms,", floors:", house.floors)
	fmt.Println("Is our house on the market?:",house.isAvailable)
	fmt.Println("")
	fmt.Println("Calling information from our employee struct:")
	fmt.Println("Our new employees first and last name:",employee.firstname, employee.lastname)
	fmt.Println("Our new employees age and date of birth:",employee.age,employee.dob)
	fmt.Println("Our new employees department is",employee.department,"and his current grade is",employee.grade,".Is he working still working with us?",employee.isHired)
	fmt.Println("Our new employees salary is",employee.salary, "and his tax bracket is",employee.taxbracket)

}
