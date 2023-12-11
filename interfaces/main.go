package main

import (
	"fmt"
	"math"
)

type square struct {
	length int
}

type circle struct {
	radius int
}

// square methods

func (s square) area() int {
	return s.length * s.length
}

func (s square) circumfrence() int {
	return s.length * 4
}

// circle methods

func (c circle) area() int {
	return int(math.Pi * float64(c.radius) * float64(c.radius))
}

func (c circle) circumfrence() int {
	return int(math.Pi * 2 * float64(c.radius))
}

type shape interface {
	area() int
	circumfrence() int
}

var shapes = []shape{
	circle{radius: 7},
	square{length: 12},
	circle{radius: 90},
	square{length: 10},
}

func main() {
	for _, shape := range shapes {
		fmt.Printf("Shape is of type: %T,Area is: %v Perimeter is: %v \n", shape, shape.area(), shape.circumfrence())
	}
}
