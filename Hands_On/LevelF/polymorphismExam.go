package main

import "fmt"

type lord struct {
	firstName	 string
	lastName	 string
	age 		 int
	house		string
	creature	string
	isALord		bool
	motto		string
}

type mother struct {
	lord
	isAGod	 bool
}

type rulers interface {
	speaking()
}

func main() { //nolint:typecheck

	house1 := lord{
		firstName: "Karl",
		lastName:  "Heisenberg",
		age:       100,
		house:     "Heisenberg",
		creature:  "Humanoid",
		isALord:   true,
		motto:     "'The strong will destroy the weak, that's the way of the world'",
	}
	house2 := lord{
		firstName: "Alcina",
		lastName:  "Dimitrescu",
		age:       107,
		house:     "Dimitrescu",
		creature:  "Vampire",
		isALord:   true,
		motto:     "'You stupid Man-thing!!!!'",
	}
	house3 := lord{
		firstName: "Salvatore",
		lastName:  "Moreau",
		age:       100,
		house:     "Moreau",
		creature:  "Fish-like creature",
		isALord:   true,
		motto:     "'Oh, Mother Miranda, if it's for you I'll do anything'",
	}

	house4 := lord{
		firstName: "Donna",
		lastName:  "Beneviento",
		age:       100,
		house:     "Beneviento",
		creature:  "Humanoid",
		isALord:   true,
		motto:     "'Did you like seeing your dead wife !?!'",
	}

	miranda := mother {
		  lord: lord{
		  	firstName: "Miranda",
		  	lastName:  "N/A",
		  	age:		200,
		  	house:		"all four houses",
		  	creature: 	"Humanoid",
		  	isALord: 	false,
		  	motto:      "'Remember from whence you came'",
				},
				isAGod: true,
	}

	allRulers(house1)
	allRulers(house2)
	allRulers(house3)
	allRulers(house4)

	allRulers(miranda)

}

func (l lord) speaking(){
	fmt.Println(l.firstName,"of house",l.house,"says",l.motto)
}

func (m mother) speaking(){
	fmt.Println(m.firstName,"is the ruler of",m.house,". She says",m.motto)
}

func allRulers(r rulers){
	r.speaking()
}