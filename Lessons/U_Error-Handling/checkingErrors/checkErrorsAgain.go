package main

import "fmt"

func main() {
	var ans1, ans2, ans3 string
	var ans4 float64

	fmt.Print("First Name: ")
	_, err := fmt.Scan(&ans1)
	if err != nil {
		panic(err) // The "panic" function stops the normal execution
	}
	fmt.Print("Last Name: ")
	_, err = fmt.Scan(&ans2) // here we are just reassigning our "err" variable with "="
	if err != nil {
		panic(err)
	}
	fmt.Print("Job Title: ")
	_, err = fmt.Scan(&ans3)
	if err != nil {
		panic(err)
	}
	fmt.Print("Gross Annual Salary: ")
	_, err = fmt.Scan(&ans4)
	if err != nil {
		panic(err)
	}

	fmt.Println(ans1, ans2, ans3, ans4)

}
