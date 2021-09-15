package main

import "fmt"

type monster interface {
	speak() string
}

func printDetails(m monster) {
	fmt.Println(m.speak())
}

type lord struct {
	Name    string
	Age     int
	IsHuman bool
	Motto   string
}

type ruler struct {
	lord
	isRuler bool
}

func (l lord) speak() string {
	return l.Name + " says " + l.Motto
}

func (r *ruler) speak() string {
	return r.Name + " says " + r.Motto
}

func main() {
	dimitrescu := lord{
		Name:    "Alcina Dimistrescu",
		Age:     200,
		IsHuman: false,
		Motto:   "Rest while you can, because I will find you and I will break you!!!",
	}
	heisenburg := lord{
		Name:    "Karl Heisenburg",
		Age:     100,
		IsHuman: false,
		Motto:   "The strong will destroy the weak, that’s the way of the world.",
	}
	moreau := lord{
		Name:    "Salvatore Moreau",
		Age:     100,
		IsHuman: false,
		Motto:   "Oh, Mother Miranda... If it's for you, I'd do anything",
	}
	beneviento := lord{
		Name:    "Donna Beneviento",
		Age:     110,
		IsHuman: false,
		Motto:   "Don't leave...I can't let you..",
	}

	miranda := &ruler{
		lord: lord{
			Name:    "Mother Miranda",
			Age:     200,
			IsHuman: false,
			Motto:   "Do not forget from whence you came",
		},
		isRuler: true,
	}
	printDetails(dimitrescu)
	printDetails(heisenburg)
	printDetails(moreau)
	printDetails(beneviento)
	printDetails(miranda)
}
