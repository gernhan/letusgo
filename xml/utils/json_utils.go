package utils

import (
	"encoding/json"
)

func ParseJSON(o interface{}) (string, error) {
	data, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// FromJSON parses a JSON string into the provided type using TypeReference
func FromJSON(data string, v interface{}) error {
	if len(data) == 0 {
		return nil
	}
	err := json.Unmarshal([]byte(data), v)
	if err != nil {
		return err
	}
	return nil
}

// FromNonNullJSON parses a non-empty JSON string into the provided type
func FromNonNullJSON(data string, v interface{}) error {
	return json.Unmarshal([]byte(data), v)
}
