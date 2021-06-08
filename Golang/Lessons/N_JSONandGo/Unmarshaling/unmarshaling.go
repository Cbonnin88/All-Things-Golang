package main

import (
	"encoding/json"
	"fmt"
)

// France Creating our struct to hold the JSON data:
type France struct {
	Country    string `json:"country"`
	Capital    string `json:"capital"`
	Continent  string `json:"continent"`
	OfficalLan string `json:"officalLan"`
	Currency   string `json:"currency"`
	Pop        int    `json:"pop"`
	EconRank   int    `json:"econRank"`
	 Government struct { // This is not a embedded Struct but an extension to our France struct
		HeadOfState string `json:"headOfState"`
		HeadOfGov   string `json:"headOfGov"`
	} `json:"government:"`

	LargestCities    []string      `json:"largestCities"`
	MarriageEquality []interface{} `json:"marriageEquality"` // here we use 'interface' because we have two different data types, a boolean and an integer
}




func main() { //nolint:typecheck

	/* Unmarshaling is the opposite of marshaling, Unmarshaling is decoding data*/


	// This is our JSON data:
	countryData := []byte(`{ 
		"country":"France",
		"capital":"Paris",
		"officalLan": "French",
    	"currency": "Euro",
    	"pop": 67000000,
    	"econRank": 6,
    	"government:": {
      		"headOfState": "Emmanuel Macron",
      		"headOfGov": "Jean Castex"
   			 },
    	"largestCities": ["Marseille","Lyon"],
   		 "marriageEquality": [true, 2013]
		}`)

	user := France{}
	err := json.Unmarshal(countryData,&user) // Here we unmarshal the JSON data
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(user)



}
