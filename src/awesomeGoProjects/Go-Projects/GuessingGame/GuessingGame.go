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

	for {
		fmt.Println("Choose a number between one and ten")
		fmt.Scan(&guess)
		if guess == randomNum {
			fmt.Printf("Great guess, it took you %d tries", try)
			break
		} else if guess > randomNum {
			fmt.Println("Too High")
		} else if guess < randomNum {
			fmt.Println("Too Low")
		}
		try++
	}
}
