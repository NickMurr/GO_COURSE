package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

// Person interface
type Person struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     string `json:"age"`
}

// 1. Marshall
func ex1() {
	p1 := Person{
		Name:    "Nick",
		Surname: "Muraviov",
		Age:     "23",
	}

	p2 := Person{
		Name:    "Ben",
		Surname: "Award",
		Age:     "27",
	}
	// people := p1
	people := []Person{p1, p2}

	bs, err := json.Marshal(people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
}

// 2. Unmarshall

type personEx2 struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     string `json:"age"`
}

func ex2() {
	s := `[{"name":"Nick","surname":"Muraviov","age":"23"},{"name":"Ben","surname":"Award","age":"27"}]`
	fmt.Println(s)
	var people []personEx2

	err := json.Unmarshal([]byte(s), &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(people)

	for i, person := range people {
		fmt.Println("Person #", i+1, person.Name, person.Surname)

	}

}

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func ex3() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	// fmt.Println(users)

	// your code goes here

	out := json.NewEncoder(os.Stdout)
	err := out.Encode(users)
	if err != nil {
		fmt.Println("We did something wrong: ", err)
	}
}

// sort
func ex4() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	// sort xi
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}

type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByLast []user

func (a ByLast) Len() int           { return len(a) }
func (a ByLast) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLast) Less(i, j int) bool { return a[i].Last < a[j].Last }

func ex5() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}
	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		sort.Strings(u.Sayings)
		for _, v := range u.Sayings {
			fmt.Println("\t", v)
		}
	}

	fmt.Println(users)
	fmt.Println("------------------")

	sort.Sort(ByLast(users))
	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		for _, v := range u.Sayings {
			fmt.Println("\t", v)
		}
	}
	fmt.Println("------------------")

	sort.Sort(ByAge(users))
	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		for _, v := range u.Sayings {
			fmt.Println("\t", v)
		}
	}

	// your code goes here
}
