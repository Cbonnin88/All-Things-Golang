package main

import "fmt"

func main() {

	// Sprint format
	grade := 100
	compliment := "Great job!"
	teacherSays := fmt.Sprint("You scored a ", grade, " on the test!", compliment)
	fmt.Println(teacherSays)

	// Sprintln format
	salary := 35000
	quote := "house"
	coachSays := fmt.Sprintln("So you earn about", salary, "euros annually?,. No wonder you have a nice", quote)
	fmt.Print(coachSays)

	// If we need to interpolate a string, without printing it, then we use "fmt.Sprintf()"

	city := "Chartres"
	move1 := fmt.Sprintf("We are moving to the city of %v\n", city)
	fmt.Print(move1)

	// anther way of writing the Sprint methods (valid for Sprint, Spirintln and Sprintf):
	stepA := "Breath in...."
	stepB := "Breath out..."
	meditation := fmt.Sprintln(stepA, stepB)
	fmt.Print(meditation)
}
