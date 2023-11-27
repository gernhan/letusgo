package main

import (
	"fmt"
	"time"
)

type S struct {
	Name string
}

func main() {

	doneCh := make(chan struct{})
	go func() {
		defer close(doneCh)
		for i := 0; i < 10; i++ {
			fmt.Println(i)
			<-time.After(time.Second)
		}
	}()

	go func() {
		<-time.After(5 * time.Second)
		var s *S
		fmt.Println(s.Name)
	}()

	<-doneCh
}