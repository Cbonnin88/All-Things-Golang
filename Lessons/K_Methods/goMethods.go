package main

import "fmt"

type person struct {
	firstName	string
	lastName	string
	age 		int
	dob			string
}

type employee struct {
	person
	isEmployed	bool
	title		string
	grade		int
	department	string
	salary		int
}

func main() {
	job1 := employee{
			person: person{
				"Christopher",
				"Bonnin",
				32,
				"10/18/1988",
			},
			isEmployed: true,
			title: "Go Developer",
			grade: 1,
			department: "IT/Tech",
			salary: 28000,
	}

	job2 := employee{
			person: person{
				firstName: "Sandra",
				lastName:  "Clark",
				age: 40,
				dob: "05/24/1980",
			},
			isEmployed: true,
			title: "Head Scala Developer",
			grade: 10,
			department: "IT/Tech",
			salary: 70000,
	}
	fmt.Println(job1)
	job1.work()
	job2.work()
}
/*

A method is a function with a receiver
structure: func (r receiver) identifier(parameters) (return(s)) {
				code...
}
*/

// creating our function and using a switch statement
func (e employee) work() {
	switch {
	case e.grade <= 2:
		level := "I am an entry level developer "
		fmt.Println("My name is", e.firstName, "and I am a", e.title, ". My grade is", e.grade, ",so", level+
			". I work in the", e.department, "and I earn", e.salary, "a year")
	case e.grade > 2 && e.grade <= 4:
		level := "I am a junior level developer"
		fmt.Println("My name is", e.firstName, "and I am a", e.title, ". My grade is", e.grade, ",so", level+
			". I work in the", e.department, "and I earn", e.salary, "a year")
	case e.grade > 5 && e.grade <= 7:
		level := "I am a mid level developer"
		fmt.Println("My name is", e.firstName, "and I am a", e.title, ". My grade is", e.grade, ",so", level+
			". I work in the", e.department, "and I earn", e.salary, "a year")
	case e.grade > 7 && e.grade <= 9:
		level := "I am a senior level developer"
		fmt.Println("My name is", e.firstName, "and I am a", e.title, ". My grade is", e.grade, ",so", level+
			". I work in the", e.department, "and I earn", e.salary, "a year")
	case e.grade == 10:
		level := "I am a head developer"
		fmt.Println("My name is", e.firstName, "and I am a", e.title, ". My grade is", e.grade, ",so", level+
			". I work in the", e.department, "and I earn", e.salary, "a year")
	}
}




