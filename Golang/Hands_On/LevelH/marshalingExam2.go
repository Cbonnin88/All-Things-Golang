package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	Firstname	string
	Lastname	string
	Age 		int
	Ishuman		bool
	FavColor	[]string
}

func main() {

	user1 := user{
		Firstname: "Thomas",
		Lastname:  "Wellington",
		Age: 33,
		Ishuman: true,
		FavColor: []string{"Blue","Green"},
	}

	user2 := user{
		Firstname: "Gina",
		Lastname: "Arenberg",
		Age: 18,
		Ishuman: false,
		FavColor: []string{"Red","Orange"},
	}

	user3 := user{
		Firstname: "Lionel",
		Lastname: "Duchateau",
		Age: 65,
		Ishuman: false,
		FavColor: []string{"Black","Brown"},
	}

	user4 := user{
		Firstname: "Lilly",
		Lastname: "Fields",
		Age: 40,
		Ishuman: true,
		FavColor: []string{"Teal","Azule"},
	}

	people := []user{user1,user2,user3,user4}
	 fmt.Println(people)
	fmt.Println("")


	err := json.NewEncoder(os.Stdout).Encode(people)
	if err != nil {
		fmt.Println("Uh oh, something is amiss here gurl and this is what it's saying:",err)
	}

}
