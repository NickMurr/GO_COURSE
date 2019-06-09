package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p1 := person{
		name: "Nick",
		age:  23,
	}
	fmt.Println("person befor", p1)
	changeMe(&p1)
	fmt.Println("person after", p1)
}

func changeMe(p *person) {
	p.name = "Nikita" // or
	(*p).age = 42
}
