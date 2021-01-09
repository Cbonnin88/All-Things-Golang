package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(" Go calculator")
	cmd := readLine("Enter command: [+]add, [-]subtract, [*]multiply, [/]divide: ")
	fmt.Println(cmd)
	switch cmd {
	case "+":
		num1, num2 := getUserNumbers()
		results := num1 + num2
		sResults := fmt.Sprintf("%d", results)
		fmt.Println(sResults)
	case "-":
		num1, num2 := getUserNumbers()
		results := num1 - num2
		sResults := fmt.Sprintf("%d", results)
		fmt.Println(sResults)
	case "*":
		num1, num2 := getUserNumbers()
		results := num1 * num2
		sResults := fmt.Sprintf("%d", results)
		fmt.Println(sResults)
	case "/":
		num1, num2 := getUserNumbers()
		results := float32(num1) / float32(num2)
		sResults := fmt.Sprintf("%f", results)
		fmt.Println(sResults)
	default:
		fmt.Println("Invalid input")
	}
}

func readLine(message string) string {
	fmt.Print(message)
	var input string
	fmt.Scanln(&input)
	return input
}

func getUserNumbers() (int, int) {
	num1String := readLine("First Number: ")
	num1, err := strconv.Atoi(num1String)
	if err != nil {
		fmt.Println(err)
	}
	num2String := readLine("Second Number: ")
	num2, err := strconv.Atoi(num2String)
	if err != nil {
		fmt.Println(err)
	}
	return num1, num2
}
