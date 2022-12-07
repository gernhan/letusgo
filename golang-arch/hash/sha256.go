package hash

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
)

func testSha256ForAFile(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(f)

	h := sha256.New()
	_, err = io.Copy(h, f)
	if err != nil {
		log.Fatalln("couldn't io.Copy", err)
	}

	fmt.Printf("Here's the type before Sum: %T", h)
	fmt.Printf("%v\n", h)
	xb := h.Sum(nil)
	fmt.Printf("Here's the type after Sum: %T", xb)
	fmt.Printf("%x\n", xb)

}
