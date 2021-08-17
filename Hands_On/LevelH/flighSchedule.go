package main

import (
	"encoding/json"
	"fmt"
)

type Flight struct {
	FlightNum	 int    `json:"flight-num"`
	Origin       string `json:"origin"`
	Destination  string `json:"destination"`
	Price		 int	`json:"price"`
	Layover		 bool	`json:"layover"`
	Ontime		 bool	`json:"ontime"`
}

func main() {
	flightOne := Flight{
				FlightNum: 2234,
				Origin: "JFK",
				Destination: "CDG",
				Price: 800,
				Layover: true,
				Ontime: true,
	}
	flightTwo := Flight{
				FlightNum: 1234,
				Origin: "LHR",
				Destination: "PHL",
				Price: 1200,
				Layover: false,
				Ontime: false,
	}

	b, err := json.MarshalIndent(flightOne, "", " ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))

	b, err = json.MarshalIndent(flightTwo, "", " ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))

}
