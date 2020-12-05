package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	for {
		fmt.Println("Choose a number from one to 10")
		var guess int
		fmt.Scan(&guess)
		rand.Seed(time.Now().UnixNano())
		secretNumber := rand.Intn(10)

		if guess > secretNumber {
			fmt.Println("Too high")
			continue
		} else if guess < secretNumber {
			fmt.Println("Too low")
			continue
		} else if guess == secretNumber {
			fmt.Println("Great Guess")
			break
		}
	}
}
