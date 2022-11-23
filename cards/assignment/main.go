package main

import (
	"fmt"
	"strconv"
)

func main() {
	slice := [10]int{}
	for i := 0; i < 10; i++ {
		slice[i] = i + 1
	}

	for i := range slice {
		message := strconv.Itoa(i) + " is "
		if i%2 == 0 {
			message = message + "even"
		} else {
			message = message + "odd"
		}

		fmt.Println(message)
	}
}
