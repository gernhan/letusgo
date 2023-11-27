package main

import (
	"encoding/json"
	"fmt"
	"log"
)

const input = `
{
	"type": "sound",
	"description": "dynamite",
	"authority": "the Bruce Dickinson"
}
`

type Envelope struct {
	Type string
}

type Sound struct {
	Description string
	Authority   string
}

func main() {
	var env Envelope
	buf := []byte(input)
	if err := json.Unmarshal(buf, &env); err != nil {
		log.Fatal(err)
	}
	switch env.Type {
	case "sound":
		var s struct {
			Envelope
			Sound
		}
		if err := json.Unmarshal(buf, &s); err != nil {
			log.Fatal(err)
		}
		var desc string = s.Description
		fmt.Println(desc)
	default:
		log.Fatalf("unknown message type: %q", env.Type)
	}
}
