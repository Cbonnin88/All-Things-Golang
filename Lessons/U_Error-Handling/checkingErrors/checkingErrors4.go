package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Open("pie.txt") // Here we are opening our file
	if err != nil {
		fmt.Println(err) // Once again we print out our error
		return
	}

	defer file.Close() // Here we aren't handling the defer statement

	bs, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err) // once again we are printing out our error, if it exists
	}

	fmt.Println(string(bs))
}
