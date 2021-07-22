package main

import (
	"encoding/json"
	"fmt"
)

type car struct {
	Model 	string
	Year    int
	Price   float64
}

func main() {

	car1 := car {
		Model: "Ford",
		Year:   2000,
		Price:  8000,
	}

	car2 := car {
		Model:  "Audi",
		Year:   2021,
		Price:  25000,
	}

	car3 := car {
		Model: "Lexus",
		Year: 	2016,
		Price:  18000,
	}

	buyer := []car{car1,car2,car3}
	fmt.Println("Our car selection:",buyer)
	fmt.Println("")

	byteSize, err := json.MarshalIndent(buyer,"","  ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(byteSize))
}
