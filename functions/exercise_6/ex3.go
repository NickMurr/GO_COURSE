package main

import "fmt"

func ex3() {
	defer foo31()
	foo3()
}

func foo3() {
	fmt.Println("Foo 1 simple")
}

func foo31() {
	fmt.Println("Foo 2 defer")
}
