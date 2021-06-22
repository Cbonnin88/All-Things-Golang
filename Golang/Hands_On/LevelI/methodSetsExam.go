package main

import "fmt"

type Person struct {
	Name	string
	Age 	int
	Motto	string
	IsHuman	bool
}

type Human interface {
	speak()
}

func main() { //nolint:typecheck

	p1 := Person{
		Name:    "Alcina Dimistrescu",
		Age:     100,
		Motto:   "You stupid Man-Thing",
		IsHuman: true,
	}
	saySomething(&p1)
	p1.speak()





	p2 := Person{
		Name:    "Toto",
		Age:     5,
		Motto:   "Bark Bark",
		IsHuman: true,
	}
	saySomething(&p2)
	p2.speak()




}

func (p *Person) speak()  {
	fmt.Println(p.Motto)
}

func saySomething (h Human) {
	h.speak()

}



