package main

import (
	"fmt"
	. "golang-arch/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/encode", Foo)
	http.HandleFunc("/decode", Bar)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Caught error: ", err)
	}
}
