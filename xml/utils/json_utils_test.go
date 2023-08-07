package utils

import (
	"fmt"
	"testing"
)

func TestAll(t *testing.T) {
	// Example usage of the JsonUtils struct
	// Create some example data
	data := map[string]interface{}{
		"foo": "bar",
		"baz": 123,
	}

	// Convert to a pretty-printed JSON string
	jsonString, err := ParseJSON(data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the JSON string
	fmt.Println(jsonString)

	// Example usage of FromJSON function
	type MyData struct {
		Foo string `json:"foo"`
		Baz int    `json:"baz"`
	}

	jsonString = `{"foo": "hello", "baz": 456}`
	var myData MyData

	err = FromJSON(jsonString, &myData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the parsed data
	fmt.Printf("Parsed Data: %+v\n", myData)
}
