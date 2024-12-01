package utils

import (
	"log"
	"os"
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
