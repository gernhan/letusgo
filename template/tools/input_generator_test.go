package tools

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestGenerateString(t *testing.T) {
	inputGenerator := NewInputGenerator()

	// Test that the generated string has a valid length.
	maxLength := 100
	generatedString := inputGenerator.GenerateString(maxLength)
	if len(generatedString) > maxLength {
		t.Errorf("Generated string length is greater than %d", maxLength)
	}

	// Test that the generated string only contains lowercase letters.
	for _, char := range generatedString {
		if char < 'a' || char > 'z' {
			t.Errorf("Generated string contains invalid character: %c", char)
		}
	}
}

func TestCreateTestFile(t *testing.T) {
	inputGenerator := NewInputGenerator()

	outputFile := "test_input.txt"
	noLines := 20

	err := inputGenerator.CreateTestFile(outputFile, noLines)
	if err != nil {
		t.Errorf("Error creating test file: %v", err)
	}

	// Read the contents of the created test file.
	content, err := readFileContent(outputFile)
	if err != nil {
		t.Errorf("Error reading test file content: %v", err)
	}

	// Test that the number of lines in the test file is equal to noLines.
	lines := strings.Split(strings.TrimSpace(content), "\n")
	if len(lines) != noLines {
		t.Errorf("Unexpected number of lines in the test file. Expected: %d, Got: %d", noLines, len(lines))
	}
}

func readFileContent(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
