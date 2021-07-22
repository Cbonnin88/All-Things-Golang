package main

import "fmt"

func main() { //nolint:typecheck

// ** Assertions are for interface types **
	var country interface{} = "France"
	str, ok := country.(string) // here we assert our that our type is a string
	if ok {
		fmt.Printf("The value is of type %T\n",str)
	} else {
		fmt.Printf("value is not a string\n")
	}

	var city interface{} = 23
	num, ok := city.(float32) // here we assert our type
	if ok {
		fmt.Printf("The value is of type %T\n",num)
	} else {
		fmt.Printf("the value is not an int\n")
	}

		// Type assertion with a panic error:
		var floater interface {}= 198.88 // float64
		num2, ok := floater.(float32)
		if !ok {
			err := fmt.Errorf("invalid type: data if of type %T", num2)
			fmt.Printf(err.Error())
		} else {
			fmt.Printf("Our value is of type %T",num2)
		}
}
