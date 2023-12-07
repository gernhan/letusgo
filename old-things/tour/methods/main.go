package main

import (
	"fmt"
	"math"
)

/**
There are two reasons to use a pointer receiver.

The first is so that the method can modify the value that its receiver points to.

The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.
 */

type Vertex struct {
	X, Y float64
}

// Remember: a method is just a function with a receiver argument.
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) NotScale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Here's Abs written as a regular function with no change in functionality.
func Abs2(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(Abs2(v))

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	v.Scale(10)
	fmt.Println(v.Abs())

	p := &v

	p.Scale(10)
	fmt.Println(v.Abs())

	v.NotScale(10)
	fmt.Println(v.Abs())
}