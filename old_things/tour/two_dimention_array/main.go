package main

import (
	"fmt"
)

func Pic(dx, dy int) [][]uint8 {
	var returnValue [][]uint8
	for i := 0; i < dx; i++ {
		returnValue = append(returnValue, make([]uint8, dy))
	}
	return returnValue

}

func main() {
	fmt.Println(Pic(2,3))
}
