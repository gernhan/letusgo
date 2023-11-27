package main

import (
	"fmt"
	"strings"
)

type Vertex struct {
	Lat, Long float64
}

var literal = map[string]Vertex{
	"Google": {
		37.42202, -122.08408,
	},
}

func WordCount(s string) map[string]int {
	stringArray := strings.Split(s, " ")
	toReturn := make(map[string]int)
	for i := range stringArray {
		v, ok := toReturn[stringArray[i]]
		if ok {
			toReturn[stringArray[i]] = v + 1
		} else {
			toReturn[stringArray[i]] = 1
		}
	}
	return toReturn
}

func main() {
	n := make(map[string]Vertex)
	n["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(n)

	// literal
	fmt.Println(literal)

	// CRUD
	mWithCrud := make(map[string]int)

	mWithCrud["Answer"] = 42
	fmt.Println("The value:", mWithCrud["Answer"])

	mWithCrud["Answer"] = 48
	fmt.Println("The value:", mWithCrud["Answer"])

	delete(mWithCrud, "Answer")
	fmt.Println("The value:", mWithCrud["Answer"])

	v, ok := mWithCrud["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
