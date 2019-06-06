package main

import "fmt"

func main() {
	fmt.Println("Hello everyone, this is the most great and learning the GO programming language")
	foo(1)
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		} else {
			foo(i)
		}
	}
}

func foo(i int) {
	fmt.Println("I'm in foo", i)
}
