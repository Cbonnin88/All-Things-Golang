package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)


func main() {
	pwd := `catfood1989`
	bs, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pwd)
	fmt.Println(bs)
	fmt.Println("")

	loginPwd := `catfood1989`
	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPwd))
	if err != nil {
		fmt.Println("INCORRECT PASSWORD")
		return
	}
	fmt.Println("You have successfully logged in")
}
