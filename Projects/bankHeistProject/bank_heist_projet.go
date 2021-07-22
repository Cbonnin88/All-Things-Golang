package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	isHeistOn := true
	rand.Seed(time.Now().UnixNano())
	eludedGuards := rand.Intn(100)

	if eludedGuards >= 50 {
		fmt.Println("So far we've made it past the guards, let's continue.")
	} else {
		isHeistOn = false
		fmt.Println("Better luck next time")
	}

	openedVault := rand.Intn(100)

	if isHeistOn && openedVault >= 70 {
		fmt.Println("Get the money !!!")
	} else if !isHeistOn {
		fmt.Println("We can't open the vault")
	}

	leftSaftely := rand.Intn(5)

	if isHeistOn {
		switch leftSaftely {
		case 0:
			isHeistOn = false
			fmt.Println("Crap, the old lady is an assassin !!")
		case 1:
			isHeistOn = false
			fmt.Println("Looks like the vault sealed us in")
		case 2:
			isHeistOn = false
			fmt.Println("This is a blood bank !!")
		case 3:
			isHeistOn = false
			fmt.Println("Look out, its Batman !!!")
		default:
			fmt.Println("Get the teleporter started")
		}
	}
	if isHeistOn {
		amtStolen := 10000 + rand.Intn(1000000)
		fmt.Println("Not bad, we got", amtStolen, " euros !!!")
	}
}
