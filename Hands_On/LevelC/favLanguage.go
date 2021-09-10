package main

import (
	"fmt"
	"log"
)

func main() {

	fmt.Print("enter a programming language or a word: ")
	var favLang string
	_, err := fmt.Scan(&favLang)
	if err != nil {
		log.Println("Unable to comply")
	}
	switch favLang {
	case "Python":
		fmt.Println("I like Python")
	case "JavaScript":
		fmt.Println("I like JavaScript")
	case "Scala":
		fmt.Println("I like Scala")
	case "Java":
		fmt.Println("I like Java")
	case "Go":
		fmt.Println("I like Go")
	default:
		fmt.Println("I'm not a software engineer")
	}

}
