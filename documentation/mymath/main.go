// Package mymath provides ACME in math solutions
package mymath

// Sum adds an unlimited number of values of type int
func Sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

// godoc -http=:8080
