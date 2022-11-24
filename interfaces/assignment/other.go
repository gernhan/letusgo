package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		fileName := os.Args[1:2]
		file, err := os.Open(string(fileName[0]))
		if err != nil {
			fmt.Println("Error while opening the file")
			os.Exit(1)
		}

		data := make([]byte, 99999)
		count, err := file.Read(data)
		if err != nil {
			fmt.Println("Error while reading the file")
			os.Exit(1)
		}

		fmt.Println(string(data[:count]))
	} else {
		fmt.Println("Empty file")
	}
}
