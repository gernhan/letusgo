package main

import (
	"fmt"
	"time"
)

func sender(c chan int) {
	time.Sleep(10 * time.Second)
	fmt.Println("Push to the channel")
	close(c)
}

func main() {
	c := make(chan int)
	go sender(c)
	v, isOk := <- c
	if isOk {
		fmt.Println(v)
	} else {
		fmt.Println("Channel is closed")
	}
	fmt.Printf("The end")
}
