package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("A simple simulation of two round voting system")

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("\nEnter candidates' full name: ")
	scanner.Scan()
	fullName := scanner.Text()

	scanner2 := bufio.NewScanner(os.Stdin)
	fmt.Print("\nEnter Candidates' Political Party: ")
	scanner2.Scan()
	party := scanner2.Text()

	fmt.Print("\nEnter Candidates first round score: ")
	var firstRound int
	_ , err := fmt.Scan(&firstRound)
	if err != nil {
		return
	}

	scanner3 := bufio.NewScanner(os.Stdin)
	fmt.Print("\nEnter second candidates' full name: ")
	scanner3.Scan()
	fullName2 := scanner3.Text()

	scanner4 := bufio.NewScanner(os.Stdin)
	fmt.Print("\nEnter candidates' Political Party: ")
	scanner4.Scan()
	party2 := scanner4.Text()

	fmt.Print("\nEnter Candidates first round score: ")
	var firstRound2 int
	_, _ = fmt.Scan(&firstRound2)

	 winner := 50
	 noBallot := 5

	switch {
	case firstRound >= firstRound2 && firstRound >= noBallot && firstRound < winner:
		fmt.Printf("\n%v of %v and %v of %v will go on to the second round", fullName, party, fullName2, party2)
	case firstRound < noBallot && firstRound2 < noBallot:
		fmt.Printf("\n%v of %v and %v of %v did not reach the five percent ballot threshold", fullName, party, fullName2, party2)
	case firstRound >= winner:
		fmt.Printf("\n%v of %v has won the election", fullName, party)
	case firstRound2 >= winner:
		fmt.Printf("\n%v of %v has won the election", fullName2, party2)
	case firstRound2 >= firstRound && firstRound2 >= noBallot && firstRound2< winner:
		fmt.Printf("\n%v of %v and %v of %v will go on to the second round",
			fullName2,party2,fullName,party)
	}
}
