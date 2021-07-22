package main

import "fmt"

/* Logical Operators in Go:

== Equal
!= NOT Equal
> Greater Than
>= Greater Than or equal too
< Less Than
<= Less Than or equal too

&& = AND
|| = OR
! = NOT

AND:
true && true = true
true && false = false
false && true = false
false && false = false

OR:
true || true = true
true || false = true
false || true = true
false || false = false
*/

func main() {
	// AND &&:
	currentBankBalance := 4559
	budgetLimit := 1200

	if currentBankBalance >= 1100 && budgetLimit >= 1100 {
		fmt.Println("You have a positive account balance")
	} else {
		fmt.Println("You have a negative account balance")
	}

	// OR ||:

	currentSalary := 30000
	currentBonuspoints := 20000

	if currentSalary >= 100000 || currentBonuspoints >= 10000 {
		fmt.Println("Congratulations, you are eligible for our gold star pass")

	} else {
		fmt.Println("We regret to inform you that you do not meet the eligibility requirements")
	}

	// the not or ! operator negates (reverses) the value of a boolean:
	// !true = false
	// !false = true

	rich := true
	fmt.Println("This reverse our 'true' to", !rich)
	poor := false
	fmt.Println("This reverses our 'false' to", !poor)

	if !rich {
		fmt.Println("I don't need money")
	} else {
		fmt.Println("I hate being poor")
	}

	if !poor {
		fmt.Println("I need more money")
	} else {
		fmt.Println("I'm living pretty well")
	}
}
