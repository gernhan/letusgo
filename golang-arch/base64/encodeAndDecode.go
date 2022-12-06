package base64

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	"testing"
)

func TestEnDecode(t *testing.T) {
	msg := "This is totally fun get hands-on and learning it from the ground up. Thank you for sharing this info with me!"
	password := "ILoveDogs"

	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Panic("Couldn't bcrypt password", err)
	}
	bs = bs[:16]

	writer := &bytes.Buffer{}
	encWriter, err := encryptWriter(writer, bs)
	if err != nil {
		return
	}

	_, err = io.WriteString(encWriter, msg)
	if err != nil {
		log.Fatalln(err)
	}

	encrypted := writer.String()
	fmt.Println("before base64", encrypted)

	result2, err := enDecode(bs, encrypted)
	if err != nil {
		log.Panic("Couldn't enDecode", err)
	}
	fmt.Println(string(result2))
}

func enDecode(key []byte, input string) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("couldn't generate new cipher %w", err)
	}
	// initialization of vector
	iv := make([]byte, aes.BlockSize)

	stream := cipher.NewCTR(block, iv)
	buff := &bytes.Buffer{}
	streamWriter := cipher.StreamWriter{
		S: stream,
		W: buff,
	}

	_, err = streamWriter.Write([]byte(input))
	if err != nil {
		return nil, fmt.Errorf("couldn't streamWriter.Write to streamWriter %w", err)
	}
	return buff.Bytes(), nil
}

func encryptWriter(writer io.Writer, key []byte) (io.Writer, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("couldn't generate new cipher %w", err)
	}
	// initialization of vector
	iv := make([]byte, aes.BlockSize)

	stream := cipher.NewCTR(block, iv)
	return cipher.StreamWriter{
		S: stream,
		W: writer,
	}, nil
}
