package main

import "fmt"

// Sprint and Sprintln format Strings
/* Example:
grade := 100
compliment := "Great job!
teacherSays := fmt.Sprint("You scored a", grade, "on the test!", compliment
Prints : You scored a 100 on the test! Great Job!*/

func main() {
	weight := 350
	quote := "fat ass!!!"
	coachSays := fmt.Sprintln("No wonder you are", weight, "pounds!!!!. Gravvy is not a drink", quote)
	fmt.Print(coachSays)

	// If we need to interpolate a string, without printing it, then we use "fmt.Sprintf()"

	city := "Chartres"
	move1 := fmt.Sprintf("We are moving to the city of %v\n", city)
	fmt.Print(move1)

	// anther way of writing the Sprint methods (valid for Sprint, Spirintln and Sprintf):

	stepA := "fart in ..."
	stepB := "queef out..."
	meditation := fmt.Sprintln(stepA, stepB)
	fmt.Print(meditation)

}
