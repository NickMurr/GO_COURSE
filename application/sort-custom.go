package main

import (
	"fmt"
	"sort"
)

// ByName type
type ByName []person

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].First < a[j].First }

func sortCutom() {

	p1 := person{"James", "Moore", 32}
	p2 := person{"Monney", "Penny", 22}
	p3 := person{"Rick", "Austin", 17}
	p4 := person{"Gary", "Moore", 46}

	people := []person{p1, p2, p3, p4}

	fmt.Println(people)
	sort.Sort(ByName(people))
	fmt.Println(people)

}

// Application done
