package main

import (
	"fmt"
	"time"
)

type StringData struct {
	value string
}

type T struct {
	data interface{} `json:"-"`
}

func main() {
	stringObject := T{}

	stringObject.data = StringData{value: "nhannd"}

	fmt.Println(stringObject)

	// 14/12/2020
	fmt.Println(time.Now().Format("2006-01-02"))
	fmt.Println(time.Now())
}