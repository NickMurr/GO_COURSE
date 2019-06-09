package main

import "fmt"

func increment() {
	a := incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(" ")
	fmt.Println(b())
	fmt.Println(b())
}

func incrementor() func() int {
	var x int

	return func() int {
		x++
		return x
	}
}
