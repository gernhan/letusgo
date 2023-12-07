package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	upperChars  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerChars  = "abcdefghijklmnopqrstuvwxyz"
	digitChars  = "0123456789"
	specialChars = "!@#$%^&*()-_=+[]{}|;:'\",.<>/?"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run password_generator.go <password_length>")
		os.Exit(1)
	}

	length := parseInput(os.Args[1])

	if length <= 0 {
		fmt.Println("Password length must be a positive integer")
		os.Exit(1)
	}

	password := generatePassword(length)
	fmt.Println("Generated Password:", password)
}

func parseInput(characters string) int {
	length, err := strconv.Atoi(characters)
	if err != nil {
		fmt.Println("Invalid input. Please provide a valid positive integer for password length.")
		os.Exit(1)
	}
	return length
}

func generatePassword(length int) string {
	rand.Seed(time.Now().UnixNano())

	allChars := upperChars + lowerChars + digitChars + specialChars
	password := make([]byte, length)

	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(allChars))
		password[i] = allChars[randomIndex]
	}

	return string(password)
}
