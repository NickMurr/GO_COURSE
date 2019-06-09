package main

import "fmt"

func ex1() {
	x := foo1()
	y, z := bar1()
	fmt.Println(x)
	fmt.Println(y, z)
}

func foo1() int {
	return 1
}

func bar1() (int, string) {
	return 2, "Hello world"
}
