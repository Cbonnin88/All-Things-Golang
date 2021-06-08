package main

import (
	"encoding/json"
	"fmt"
)

type SensorReading struct {
	Model       string  `json:"name"`
	Year        int     `json:"year"`
	Color       string  `json:"color"`
	Tva         float32 `json:"tva"`
	IsFourDoor  bool    `json:"is_four_door"`
	IsManual    bool    `json:"is_manual"`
	Information Info    `json:"info"`
}

type Info struct {
	Description  string  `json:"description"`
	ForSale		 bool	 `json:"for_sale"`
	Country		 string	 `json:"country"`
}
func main() { //nolint:typecheck
	jsonString := `{"model":"Audi", "year": 2020, "color":"blue", "tva": 20.0, "isfourdoor": true, "isManual": false, "info":{
			"description":"A 2020 brand new Audi A6",
			"forsale": true,
			"country": "Germany"
	}}`

	var reading SensorReading
	err := json.Unmarshal([]byte(jsonString), &reading)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", reading)


}
