package main

import (
	"encoding/json"
	"fmt"
)

type Finland struct {
	Country     string `json:"country"`
	Capital     string `json:"capital"`
	Continent   string `json:"continent"`
	OfficialLan string `json:"officialLan"`
	Currency    string `json:"currency"`
	Pop         int    `json:"pop"`
	EconRank    int    `json:"econRank"`
	Government  struct {
		HeadOfState string `json:"headOfState"`
		HeadOfGov   string `json:"headOfGov"`
	} `json:"government"`
	LargestCities    []string      `json:"largestCities"`
	MarriageEquality []interface{} `json:"marriageEquality"`
}

func main() {
	countryData2 := `[{
    "country": "Finland",
    "capital": "Helsinki",
    "continent": "Europe",
    "officialLan": "Finnish",
    "currency": "Euro",
    "pop": 5536146,
    "econRank": 43,
    "government": {
      "headOfState": "Sauli Niinisto",
      "headOfGov": "Sanna Marin"
    },
    "largestCities": ["Espoo","Tampere"],
    "marriageEquality": [true, 2017]
  	}]`
	fmt.Println("Our JSON data before we unmarshal:")
	fmt.Println(countryData2)
	fmt.Println("")

	var country []Finland

	err := json.Unmarshal([]byte(countryData2),&country)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Our 'Finland struct:'",country)
	fmt.Println("")

	for i, finland := range country {
		fmt.Println("Country number:",i)
		fmt.Println("\t", finland.Country,"\n\t",finland.Capital,"\n\t",finland.OfficialLan)
		fmt.Println("Government:\n\t",finland.Government.HeadOfState,"\n\t",finland.Government.HeadOfGov)
		fmt.Println("Marriage Equality:\n\t",finland.MarriageEquality)
	}




}
