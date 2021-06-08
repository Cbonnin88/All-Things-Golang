package main

import (
	"encoding/json"
	"fmt"
)

type House struct {
	Year	int 	`json:"year"`
	Model 	string	`json:"model"`
	IsNew	bool	`json:"is_new"`
	TVA     float64  `json:"tva"`
	Buyer	Buyer	 `json:"buyer"`

}

type Buyer struct {
	IsHuman		bool	`json:"is_human"`
	Firstname	string  `json:"first_name"`
	Lastname	string	`json:"last_name"`
	Age 		int		`json:"age"`
}

func main() { //nolint:typecheck
	houseData := `{"year":2010, "model":"Victorian","isNew":true,"tva":25.0}`

	var data House
	err := json.Unmarshal([]byte(houseData),&data)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Our JSON converted into a Go Struct: %+v\n",data) // this will show us the fields
	fmt.Println("")

	// Another way of unmarshaling:
	buyData := `[{"year":2019, "model":"Tudor","is_new":false, "tva":25.0, "buyer":{
					"is_human":true, "first_name":"Alexander", "last_name":"Martin", "age": 33 } }]`

	bs := []byte(buyData)
	fmt.Printf("Our data type is %T\n",buyData)
	fmt.Printf("our byte type is %T\n",bs)
	fmt.Println("")
	var house []House

	err2 := json.Unmarshal(bs,&house)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println("from JSON to an embedded Go Struct:",house)
	fmt.Println("")

	// Getting the range and value:
	fmt.Println("Getting our range and values:")
	for i, v := range house {
		fmt.Println("This is the range:",i)
		fmt.Println("v.Year,v.Buyer.IsHuman, v.Buyer.Firstname gives us the values of:",v.Year,v.Buyer.IsHuman, v.Buyer.Firstname)
	}






}
