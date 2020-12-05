package main

import "fmt"

// .Scan() allows us to get user input

func main() {
	fmt.Println("Emilie, Jordan et Seb veut t'écouter dire quoi?")
	var answer string
	fmt.Scan(&answer)
	fmt.Print("\n")

	fmt.Println("jusqu'à?")
	var response string
	fmt.Scan(&response)
	fmt.Print("\n")

	fmt.Println("Piratage en cours.... 80% restant...")
	fmt.Println("\n")

	fmt.Printf("C'est fait, je les envoie %v via Facebook, WhatsApp et Facebook jusqu'à leur %v", answer, response)
}
