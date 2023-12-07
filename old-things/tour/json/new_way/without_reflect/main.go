package main

import (
	"encoding/json"
	"fmt"
	"log"
)

const input = `
{
	"type": "MessageType",
	"msg": {
		"description": "dynamite",
		"content": "the Bruce Dickinson"
	}
}
`

//go:generate jsonenums -type=Kind

type Kind int

const (
	MessageType Kind = iota
	ParagraphType
)

type Envelope struct {
	Type Kind        `json:"type"`
	Msg  interface{} `json:"msg"`
}

type Message struct {
	Description string `json:"description"`
	Content     string `json:"content"`
}

type Paragraph struct {
	Topic      string `json:"topic"`
	Content    string `json:"content"`
	Conclusion string `json:"conclusion"`
}

var kindHandlers = map[Kind]func() interface{}{
	MessageType:   func() interface{} { return &Message{} },
	ParagraphType: func() interface{} { return &Paragraph{} },
}

func main() {
	var raw json.RawMessage
	env := Envelope{
		Msg: &raw,
	}
	if err := json.Unmarshal([]byte(input), &env); err != nil {
		log.Fatal(err)
	}
	msg := kindHandlers[env.Type]()
	if err := json.Unmarshal(raw, msg); err != nil {
		log.Fatal(err)
	}

	newMessage := msg.(*Message)
	fmt.Println(*newMessage)
}
