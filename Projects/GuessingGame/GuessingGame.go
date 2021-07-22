package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(11)
	guessingGame(randomNum)
}
func guessingGame(randomNum int) {
	var guess int
	var try = 0

	myGuess:
	for {
		fmt.Println("Choose a number between one and ten")
		_, _ = fmt.Scan(&guess)
		switch {
		case guess > randomNum:
			fmt.Println("Too High")
		case guess < randomNum:
			fmt.Println("Too Low")
		default:
			if guess == randomNum{
				fmt.Printf("Great guess, it took you %d tries", try)
				break myGuess
			}
		}
		try++
	}
}
