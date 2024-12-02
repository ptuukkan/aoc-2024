package utils

import (
	"log"
	"os"
	"strings"
)

func ReadFile(fileName string) (string, error) {
	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	// Convert []byte to string
	text := string(fileContent)
	return text, nil
}

func SplitNewLines(input string) []string {
	lines := strings.Split(input, "\n")
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	return lines
}

func Abs(number int) int {
	if number < 0 {
		return -number
	}
	return number
}
