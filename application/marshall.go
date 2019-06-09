package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Age   int    `json:"age"`
}

func marshall() []byte {

	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   27,
	}
	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}

	people := []person{p1, p2}

	bs, err := json.Marshal(people)

	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(string(bs))
	return bs
}
