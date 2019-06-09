package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func ex4() {
	p1 := person{
		"Nikita",
		"Muraviov",
		23,
	}

	p1.speak()
}

func (p person) speak() {
	fmt.Printf("Hey, my name is %v %v, i'm %v years old", p.first, p.last, p.age)
}
