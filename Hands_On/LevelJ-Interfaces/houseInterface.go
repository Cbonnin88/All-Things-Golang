package main

import (
	"fmt"
)

type powerDrawer interface {
	draw()
}

func powerDrawn(p powerDrawer) {
	p.draw()
}

type appliance struct {
	name  string
	power int
}

func (a appliance) draw() {
	switch {
	case a.power <= 100:
		fmt.Printf("%v Eco-Friendly Score: %v\t\t Eco-Grade: Excellent\n", a.name, a.power)
	case a.power >= 101 && a.power <= 200:
		fmt.Printf("%v Eco-Friendly Score: %v\t\t Eco-Grade: Average\n", a.name, a.power)
	case a.power >= 201 && a.power <= 350:
		fmt.Printf("%v Eco-Friendly Score: %v\t\t Eco-Grade: Poor\n", a.name, a.power)
	case a.power >= 351:
		fmt.Printf("%v Eco-Friendly Score: %v\t\t Eco-Grade: Failure\n", a.name, a.power)
	}
}

func main() {

	appl1 := appliance{
		name:  "MacBook pro",
		power: 116,
	}

	appl2 := appliance{
		name:  "PlayStation 5",
		power: 99,
	}

	appl3 := appliance{
		name:  "SamSung TV",
		power: 350,
	}

	appl4 := appliance{
		name:  "Wifi LiveBox",
		power: 400,
	}

	powerDrawn(appl1)
	powerDrawn(appl2)
	powerDrawn(appl3)
	powerDrawn(appl4)
}
