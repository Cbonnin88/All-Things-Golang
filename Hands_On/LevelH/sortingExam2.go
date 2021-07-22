package main

import (
	"fmt"
	"sort"
)

type person struct {
	Firstname	string
	Lastname	string
	Age 		int
	Motto		[]string
}

type ByAge []person
type ByLast []person
type ByFirst []person

func main() {

	p1 := person {
		Firstname: "Lucinda",
		Lastname: "Ferln",
		Age: 78,
		Motto: []string{"We are not amused","Shake it babygurl","I'm famished"},
	}

	p2 := person {
		Firstname: "Morty",
		Lastname: "Fishbien",
		Age: 100,
		Motto: []string{"I'm in traction,","My boots are hurting,","Life is like a box of white chocolate"},
	}

	p3 := person {
		Firstname: "Alcina",
		Lastname: "Dimistrescu",
		Age: 44,
		Motto: []string{"You stupid man-thing,","You won't live long even if you run,","Well Well Well Ethan Winters"},
	}

	p4 := person {
		Firstname: "Marcus",
		Lastname: "Powell",
		Age: 33,
		Motto: []string{"here I am, ","Dost thou Comprehend?,","Your lips are moving"},
	}

	folks := []person{p1,p2,p3,p4}
	for _, p := range folks {
		fmt.Println(p.Firstname, p.Lastname, p.Age)
		for _, v := range p.Motto {
			fmt.Println("\t",v)
		}
	}
	fmt.Println("")
	fmt.Println("Sorted by age:")

	sort.Sort(ByAge(folks))
	for _, p := range folks {
		fmt.Println(p.Firstname,p.Lastname,p.Age,":")
		for _, v := range p.Motto {
			fmt.Println("\t",v)
		}
	}
	fmt.Println("")
	fmt.Println("Sorted by last name:")

	sort.Sort(ByLast(folks))
	for _, p := range folks {
		fmt.Println(p.Firstname,p.Lastname,p.Age,":")
		for _, v := range p.Motto {
			fmt.Println("\t",v)
		}
	}
	fmt.Println("")
	fmt.Println("Sorted by first name:")

	sort.Sort(ByFirst(folks))
	for _, p := range folks {
		fmt.Println(p.Firstname,p.Lastname,p.Age,":")
		for _, v := range p.Motto {
			fmt.Println("\t",v)
		}
	}





}

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func (l ByLast) Len() int           { return len(l) }
func (l ByLast) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l ByLast) Less(i, j int) bool { return l[i].Lastname < l[j].Lastname}

func (f ByFirst) Len() int           { return len(f) }
func (f ByFirst) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
func (f ByFirst) Less(i, j int) bool { return f[i].Firstname < f[j].Firstname}


