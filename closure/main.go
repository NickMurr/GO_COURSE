package main

import "fmt"

var x int

func main() {

	fmt.Println(x)
	x++
	fmt.Println(x)
	foo()
	fmt.Println(x)

	increment()
}

func foo() {
	fmt.Println("Hello")
	x++
}
