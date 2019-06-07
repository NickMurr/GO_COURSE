package main

import "fmt"

type funcarray func(slice ...int) int

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := sum(ii...)
	fmt.Println("all numbers: ", s1)

	s2 := even(sum, ii...)
	fmt.Println("even numbers: ", s2)

	s3 := odd(sum, ii...)
	fmt.Println("odd numbers: ", s3)

}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func even(f funcarray, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

func odd(f funcarray, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 != 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
