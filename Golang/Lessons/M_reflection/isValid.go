package main

import (
	"fmt"
	"reflect"
)

type newHouse struct {
	floors		int
	isNew		bool
	price		float64
}

func main() { //nolint:typecheck
	// IsValid() is used to check whether it is a value or not

	house1 := newHouse {
			floors: 2,
			isNew: true,
			price: 2000000,
	}
	fmt.Println(reflect.ValueOf(house1).IsValid())

	num := 12
	fmt.Println(reflect.ValueOf(num).IsValid())


}
