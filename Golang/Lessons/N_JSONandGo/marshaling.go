package main

import (
	"encoding/json"
	"fmt"
)

type lord struct {
	FirstName	string // our fields must be upper case in order to be marshalled
	LastName	string
	Age 		int
	House		string
	BOW			string
	IsLord		bool
	Motto		[]string
}

func main() {
	/* Marshaling is the process of transforming the memory representation of an object to a data format
	suitable for storage or transmission

	In other words: Marshaling is encoding the data*/

	h1 := lord {
		FirstName:  "Karl",
		LastName:   "Heisenburg",
		Age:	    100,
		House:	    "Heisenburg",
		BOW:		"Humanoid",
		IsLord: 	 true,
		Motto: 		[]string{"'The strong will destroy the weak, that's the way of the world'","'I'm not a child!!!'"},
	}

	h2 := lord {
		FirstName: "Alcina",
		LastName:  "Dimitrescu",
		Age:		107,
		House:      "Dimitrescu",
		BOW:		"Vampire",
		IsLord: 	 true,
		Motto:      []string{"'You stupid Man-thing!!!'", "'You won't live long even if you run!!!'"},
	}

	h3 := lord {
		FirstName: "Salvatore",
		LastName:  "Moreau",
		Age:		100,
		House:		"Moreau",
		BOW:		"Fish",
		IsLord: 	 true,
		Motto:		[]string{"'oh, Mother Miranda, if it's for you I'll do anything'", "'This is MY territory!!!'"},
	}

	h4 := lord {
		FirstName: "Donna",
		LastName:  "Beneviento",
		Age:	    100,
		House:		"Beneviento",
		BOW:		"Human",
		IsLord: 	 true,
		Motto:		[]string{"'Remember from whence you came'", "'He's aaaawwwakkke !!!!'"},
	}

	fourLords := []lord{
		h1,
		h2,
		h3,
		h4,
	}
	fmt.Println(fourLords)
	fmt.Println("")

	// Marshaling our struct:
	bs, err := json.MarshalIndent(fourLords,"","  ") // Here we indent our code for it to be more readable
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))


}
