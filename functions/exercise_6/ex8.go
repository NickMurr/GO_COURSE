package main

import "fmt"

func ex8() {
	f := foo8()
	fmt.Println(f())
	fmt.Println(foo8()())
}

func foo8() func() int {
	return func() int {
		return 42
	}
}
