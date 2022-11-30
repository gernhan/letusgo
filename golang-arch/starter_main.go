package main

import (
	"crypto/hmac"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type person struct {
	First string
}

func main() {
	fmt.Println(base64.StdEncoding.EncodeToString([]byte("user:pass")))

	pass := "123456"
	hashedPass, err := hashPassword(pass)
	if err != nil {
		panic(err)
	}

	err = comparePassword(pass, hashedPass)
	if err != nil {
		log.Fatalln("Not logged in")
	}
	log.Println("Logged in!")
}
func signMessage(msg string) (string, error) {
	hmac.New()
}

func hashPassword(password string) (string, error) {
	fromPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error while generating bcypt hash from password: %w", err)
	}
	return string(fromPassword), nil
}

func comparePassword(password string, hashedPass string) error {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hashedPass))
	if err != nil {
		return fmt.Errorf("invalid password: %w", err)
	}
	return nil
}
func foo(writer http.ResponseWriter, request *http.Request) {
	p1 := person{
		First: "Jenny",
	}

	p2 := person{
		First: "James",
	}

	people := []person{p1, p2}

	err := json.NewEncoder(writer).Encode(people)
	if err != nil {
		log.Println("Encode bad data: ", err)
	}
}

func bar(writer http.ResponseWriter, request *http.Request) {
	var p1 person
	err := json.NewDecoder(request.Body).Decode(&p1)
	if err != nil {
		log.Println("Decode bad data: ", err)
	}
}
