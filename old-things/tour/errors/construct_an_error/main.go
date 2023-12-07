package main

import (
	"fmt"
	"math"
	"time"
)

type ErrNegativeSqrt struct {
	number float64
}

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v",
		e.number)
}

func Sqrt(x float64) (float64, *ErrNegativeSqrt) {
	if x < 0 {
		return 0,  &ErrNegativeSqrt{
			x,
		}
	}
	return math.Sqrt(x), nil
}


type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
