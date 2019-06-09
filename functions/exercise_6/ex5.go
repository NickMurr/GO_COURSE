package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	length float64
}

type shape interface {
	area() float64
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (s square) area() float64 {
	return s.length * s.length
}

func info(s shape) {
	fmt.Println(s.area())
}

func ex5() {
	circ := circle{3}
	squa := square{3}

	info(circ)
	info(squa)
}
