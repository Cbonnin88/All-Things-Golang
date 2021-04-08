package main

import "fmt"

func main() {
	fmt.Print("Enter your full first name: ")
	var firstName string
	_, _ = fmt.Scan(&firstName)

	fmt.Print("Enter your full middle name: ")
	var middleName string
	_, _ = fmt.Scan(&middleName)

	fmt.Print("Enter your full last name: ")
	var lastName string
	_, _ = fmt.Scan(&lastName)

	var fullName string = firstName + " " + middleName + " " + lastName

	fmt.Println(firstName, " has a total of ", len(firstName), " letters")
	fmt.Println(middleName, " has a total of ", len(middleName), " letters")
	fmt.Println(lastName, " has a total of ", len(lastName), " letters")

	fmt.Println(fullName, " has a total of ", len(fullName), " letters")
}
