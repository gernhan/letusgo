package main

import (
	"fmt"
	"strings"
)

type S struct {
	T *T
}

type T struct {
	Name string
}

func printUpper(s S) {
	s.T.Name = strings.ToUpper(s.T.Name)
	fmt.Println(s.T.Name)
}

func main() {
	s := S{
		T: &T{
			Name: "foo",
		},
	}
	printUpper(s)
	fmt.Println(s.T.Name)
}