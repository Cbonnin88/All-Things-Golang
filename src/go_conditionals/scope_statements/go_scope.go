package main

import "fmt"

func main() {

	x := 30000
	if salary := x; salary < 27000 {
		fmt.Println(salary, "The salary offer is a bit low is there any room for negociation ?")
	} else {
		fmt.Println("The salary offer is what I was hoping, I'll take it")
	}

	switch hired := "Vampire Slayer"; hired {
	case "doctor":
		fmt.Println("Welcome to the team, Dr. Bonnin")
	case "lawyer":
		fmt.Println("Welcome to the firm, Couselor Bonnin")
	case "professeor":
		fmt.Println("Welcome to the University, Professor Bonnin")
	case "Vampire Slayer":
		fmt.Println("One girl in all the world, the chosen one")
	default:
		fmt.Println("Are you sure you're at the right interview ?")
	}

}
