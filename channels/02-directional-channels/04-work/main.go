package main

import "fmt"

func main() {
	c := make(chan int)
	cr := make(<-chan int) // receive
	cs := make(chan<- int) // send
	fmt.Println("------")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("c\t%T\n", cr)
	fmt.Printf("c\t%T\n", cs)
}
