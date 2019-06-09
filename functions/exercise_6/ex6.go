package main

import "fmt"

func ex6() {
	s := "I'm an anonymous func"
	func(s string) {
		fmt.Println(s)
	}(s)
}
