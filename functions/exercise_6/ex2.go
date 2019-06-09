package main

import "fmt"

func ex2() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	s1 := foo2(x...)
	fmt.Println(s1)

	s2 := bar2(x)
	fmt.Println(s2)

	s3 := bar22(x)
	fmt.Println(s3)
}

func foo2(n ...int) int {
	total := 0
	for _, v := range n {
		total += v
	}
	return total
}

func bar2(n []int) int {
	return foo2(n...)
}

func bar22(n []int) int {
	total := 0
	for _, v := range n {
		total += v
	}
	return total
}
