package handler

import (
	"encoding/json"
	. "golang-arch/dto"
	"log"
	"net/http"
)

func Foo(writer http.ResponseWriter, request *http.Request) {
	p1 := Person{
		First: "Jenny",
	}

	p2 := Person{
		First: "James",
	}

	people := []Person{p1, p2}

	err := json.NewEncoder(writer).Encode(people)
	if err != nil {
		log.Println("Encode bad data: ", err)
	}
}

func Bar(writer http.ResponseWriter, request *http.Request) {
	var p1 Person
	err := json.NewDecoder(request.Body).Decode(&p1)
	if err != nil {
		log.Println("Decode bad data: ", err)
	}
}
