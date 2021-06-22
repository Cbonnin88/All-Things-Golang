package main

import "fmt"

// .Scan() allows us to get user input

func main() {
	fmt.Println("Emilie, Jordan et Seb veut t'écouter dire quoi?")
	var answer string
	_, err := fmt.Scan(&answer)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print("\n")

	fmt.Println("jusqu'à?")
	var response string
	_, err2 := fmt.Scan(&response)
	if err != nil {
		fmt.Println(err2)
	}
	fmt.Print("\n")

	fmt.Println("Piratage en cours.... 80% restant...")
	fmt.Println("\n")

	fmt.Printf("C'est fait, je les envoie %v via Facebook, WhatsApp et Facebook jusqu'à leur %v", answer, response)
}
