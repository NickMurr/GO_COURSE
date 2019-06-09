package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.firstName, s.lastName, " - the agent speak")
}
func (p person) speak() {
	fmt.Println("I am", p.firstName, p.lastName, " - the person speak")
}

type human interface {
	speak()
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I was passed into barrrrrr", h.(person).firstName, h.(person).lastName)
	case secretAgent:
		fmt.Println("I was passed into barrrrrr", h.(secretAgent).firstName)

	}
	fmt.Println("i was passed into bar", h)
}

type hotdog int

func main() {
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	p1 := person{
		firstName: "Dr.",
		lastName:  "Yes",
	}

	sa1.speak()
	p1.speak()

	fmt.Println(p1)

	bar(sa1)
	bar(p1)

	var z hotdog = 432
	fmt.Printf("%T\n", z)

}
