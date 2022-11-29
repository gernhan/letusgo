package main

import (
	"fmt"
	"math"
)

type shape interface {
	getArea() float64
	printArea()
}
type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.height * t.base
}

func (t triangle) printArea() {
	fmt.Println(t.getArea())
}

func (s square) getArea() float64 {
	return 0.5 * math.Pow(s.sideLength, 2)
}

func (s square) printArea() {
	fmt.Println(s.getArea())
}
