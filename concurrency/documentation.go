package main

import "fmt"

func doSomething(x int) int {
	return x * x
}

func documentation() {

	ch := make(chan int)

	go func() {
		ch <- doSomething(5)
	}()
	fmt.Println(<-ch)

}
