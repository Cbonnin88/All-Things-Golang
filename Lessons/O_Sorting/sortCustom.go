package main

import (
	"fmt"
	"sort"
)

type person struct {
	firstName string
	age       int
	grade     float64
}

type ByAge []person
type ByName []person
type ByGrade []person
type ByOldest []person

func main() { //nolint:typecheck

	p1 := person{
		firstName: "Mark",
		age:       49,
		grade:     0.11,
	}

	p2 := person{
		firstName: "Linda",
		age:       20,
		grade:     9.1,
	}

	p3 := person{
		firstName: "Adrien",
		age:       35,
		grade:     15.11,
	}

	p4 := person{"Alcina", 18, 22.3}

	people := []person{p1, p2, p3, p4}
	fmt.Println("Unsorted:", people)

	sort.Sort(ByAge(people))
	fmt.Println("")
	fmt.Println("Names sorted by age:", people)

	sort.Sort(ByName(people))
	fmt.Println("Names sorted by letter:", people)

	sort.Sort(ByGrade(people))
	fmt.Println("Sorted by grade:", people)

	sort.Sort(ByOldest(people))
	fmt.Println("Ages sorted from oldest to youngest:", people)

}

// Len Sorting by age:
func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].age < a[j].age }

// Len Sorting by name:
func (n ByName) Len() int           { return len(n) }
func (n ByName) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n ByName) Less(i, j int) bool { return n[i].firstName < n[j].firstName }

// Len Sort by grade:
func (g ByGrade) Len() int           { return len(g) }
func (g ByGrade) Swap(i, j int)      { g[i], g[j] = g[j], g[i] }
func (g ByGrade) Less(i, j int) bool { return g[i].grade < g[j].grade }

// Len Sort by age, highest to lowest:
func (ah ByOldest) Len() int           { return len(ah) }
func (ah ByOldest) Swap(i, j int)      { ah[i], ah[j] = ah[j], ah[i] }
func (ah ByOldest) Less(i, j int) bool { return ah[i].age > ah[j].age }
