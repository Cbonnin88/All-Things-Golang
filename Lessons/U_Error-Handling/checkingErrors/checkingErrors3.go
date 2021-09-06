package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("houses.txt") // Here we creat a new file with the os package
	if err != nil {
		fmt.Println(err) // here we print out our err if there is one
		return
	}
	defer func(f *os.File) { // Here we handle the error for our defer statement, by creating a defer function
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	r := strings.NewReader("Chartres: House, 300000, 4 bedrooms\nChessy: House, 400000, 3 bedrooms") // we passed in a string into our reader
	_, err = io.Copy(f, r)
	if err != nil {
		fmt.Println(err)
	}
}
