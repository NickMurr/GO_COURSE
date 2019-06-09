package main

import (
	"fmt"
	"sort"
)

func sortSome() {
	xi := []int{10, 9, 8, 7, 6, 5, 4, 1, 2, 3}
	xs := []string{"James", "Q", "M", "Monneypenny", "Dr. No"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println("---------")
	sort.Strings(xs)
	fmt.Println(xs)
}
