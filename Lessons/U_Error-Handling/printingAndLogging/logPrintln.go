package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Create("log.txt") // here we are creating a file called log
	if err != nil {
		fmt.Println(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)
	log.SetOutput(file)

	file2, err := os.Open("cats")
	if err != nil {
		log.Println("err occurred", err) // logs our error to our log.txt that we created earlier, with a timestamp
	}
	defer func(file2 *os.File) {
		err := file2.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file2)

	fmt.Println("Check the log.txt file in your directory")

}
