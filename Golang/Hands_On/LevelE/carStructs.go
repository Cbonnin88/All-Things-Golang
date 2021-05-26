package main

import "fmt"

type  vehicle struct {
	doors 	int
	color	string
}

type truck struct {
	vehicle   vehicle
	fourwheel bool

}

type sedan struct {
	vehicle vehicle
	luxury bool
}

func main() { //nolint:typecheck

	ford := truck {
		vehicle: vehicle{
			doors: 5,
			color: "orange",
		},
		fourwheel: false,
	}

	suburu := sedan{
			vehicle: vehicle{
				doors: 6,
				color: "red",
			},
			luxury: true,
	}

	fmt.Println("our truck struct: ",ford)
	fmt.Println("Our sedan struct:",suburu)
	fmt.Println("How many doors does our truck have: ",ford.vehicle.doors)
	fmt.Println("How many doors does our sedan have: ",suburu.vehicle.doors)
	fmt.Println("Our trucks color:",ford.vehicle.color)
	fmt.Println("Our sedans color:",suburu.vehicle.color)
	fmt.Println("Does our truck have four wheels: ",ford.fourwheel)
	fmt.Println("Is our sedan a luxury sedan: ",suburu.luxury)

}
