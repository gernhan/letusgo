package tools

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type InputGenerator struct {
	random *rand.Rand
}

func NewInputGenerator() *InputGenerator {
	return &InputGenerator{
		random: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (ig *InputGenerator) GenerateString(maxLength int) string {
	length := ig.random.Intn(100) + 1
	if length > maxLength {
		length = maxLength
	}
	var array []byte
	for i := 0; i < length; i++ {
		value := ig.random.Intn(26) + 97
		array = append(array, byte(value))
	}
	return string(array)
}

func (ig *InputGenerator) CreateTestFile(outputFile string, noLines int) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	for i := 0; i < noLines; i++ {
		stringVal := ig.GenerateString(100)
		_, err := fmt.Fprintln(file, stringVal)
		if err != nil {
			return err
		}
	}

	return nil
}
