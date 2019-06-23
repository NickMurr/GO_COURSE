package main

import "fmt"

func main() {
	c := make(chan int, 2)
	c <- 41
	c <- 46
	fmt.Println(<-c)
	fmt.Println(<-c)
}
