package hash

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

const KEY = "someKey"

func testHashing() {
	var key []byte
	for i := 1; i <= 64; i++ {
		key = append(key, byte(i))
	}
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

func testSignature() {
	const msg = "This is a test for signing"
	signature, err := signMessage(msg)
	if err != nil {
		log.Fatalln("cannot generate signature")
	}
	rightPerson, err := checkSignature(msg, signature)
	if err != nil {
		return
	}
	if rightPerson {
		fmt.Println("Right person")
	} else {
		fmt.Println("Wrong person")
	}
}

func signMessage(msg string) (string, error) {
	h := hmac.New(sha512.New, []byte(KEY))
	_, err := h.Write([]byte(msg))
	if err != nil {
		return "", fmt.Errorf("error in signMessage while hashing message: %w", err)
	}
	signature := h.Sum(nil)
	return string(signature), nil
}

func checkSignature(msg string, signature string) (bool, error) {
	newSig, err := signMessage(msg)
	if err != nil {
		return false, fmt.Errorf("error in checkSignature while getting signature of message: %w", err)
	}

	same := hmac.Equal([]byte(newSig), []byte(signature))
	return same, nil
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
