package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("goose")
	if err != nil {
		log.Panicln(err)
	}
}
