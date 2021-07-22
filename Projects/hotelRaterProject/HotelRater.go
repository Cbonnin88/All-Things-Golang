package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var name string
	var idNum int         //nolint:wsl
	var userRating string //nolint:wsl

	// Front-end part:
	// use name as input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter your full name: ") 
	name, _ = reader.ReadString('\n')

	// get ID number:
	fmt.Print("Please enter your four digit identification number: ")
	_, _ = fmt.Scan(&idNum)

	// take the users rating and convert to float:
	reader = bufio.NewReader(os.Stdin)
	fmt.Print("Please rate our hotel from one to ten: ")
	userRating, _ = reader.ReadString('\n')
	myNumRating, _ := strconv.ParseFloat(strings.TrimSpace(userRating), 64)

	// Back end Part:
	fmt.Printf("Hello %v Your identification number is %v. " +
		"Thanks for giving us a %v star rating", name,idNum, myNumRating)

	defer fmt.Println("Your rating was recorded on", time.Now().Format(time.Stamp))

	switch {
	case myNumRating <= 3:
		fmt.Println("\nI'm sorry your experience wasn't great, we will shrive to do better")
	case myNumRating >= 4 && myNumRating < 6:
		fmt.Println("\nI'm glad your experience was satisfactory but we can do better")
	case myNumRating >= 6 && myNumRating < 10:
		fmt.Println("\nI'm glad you had a great experience, we are happy to hear it")
	case myNumRating == 10:
		fmt.Println("\nExcellent !!, we are always happy to hear great news !! ")
	}
}
