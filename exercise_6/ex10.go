// Closure

package main

import "fmt"

func ex10() {
	f := foo10()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func foo10() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
