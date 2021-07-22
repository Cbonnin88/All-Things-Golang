package main

import (
	"fmt"
	"reflect"
)

type vehicle struct {
	year	int
	model   string
	doors   int
	isNew   bool
}

type house struct {
	year	int 	`tag:"value1"`
	yard	bool	`tag:"value2"`
}




func main() { //nolint:typecheck
	/*
	NumField() returns number of fields in a struct
	Field(i int) returns reflect.StructField at position i in struct
	*/

	// NumField():
	car := vehicle{
		year:  2010,
		model: "Ford",
		doors:  4,
		isNew:  true,
	}

	fmt.Println("The number of fields our struct has is",reflect.ValueOf(car).NumField())
	fmt.Println("")

	// Field():
	myField := reflect.StructOf([]reflect.StructField{
		{
			Name: "Height",
			Type: reflect.TypeOf(1.82),
			Tag: `json:"height"`,
		},
		{
			Name: "Name",
			Type: reflect.TypeOf("words"),
			Tag:  `json:"name"`,
		},
		{
			Name: "Human",
			Type: reflect.TypeOf(true),
			Tag:  `json:"human"`,
		},
	})

	fmt.Println(myField.Field(0))
	fmt.Println(myField.Field(1))
	fmt.Println(myField.Field(2))
	fmt.Println("")

	// Example with NumField() and Field():
	h := house{}
	hType := reflect.TypeOf(h)

	for i := 0; i < hType.NumField(); i++ {
		myField := hType.Field(i)
		fmt.Println(myField.Name, myField.Type.Name(), myField.Tag.Get("tag"))
	}
}



