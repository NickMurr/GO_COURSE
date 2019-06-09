package main

import (
	"encoding/json"
	"fmt"
)

func unmarshall(bs []byte) {

	var people []person
	err := json.Unmarshal(bs, &people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("all of the data", people)

	for i, v := range people {
		fmt.Println("\nPERSON NUMBER -", i+1)
		fmt.Println(v.First, v.Last, v.Age)
	}
}
