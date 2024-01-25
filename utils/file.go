package utils

import (
	"encoding/json"
	"log"
	"os"
)

func SaveToFile(data interface{}, filename string) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Println("Error:", err)
		return err
	}

	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		log.Println("Error:", err)
		return err
	}

	return nil
}

func RetrieveFromFile(filename string, data interface{}) error {
	fileData, err := os.ReadFile(filename)
	if err != nil {
		log.Println("Error reading file:", err)
		return err
	}

	err = json.Unmarshal(fileData, data)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		return err
	}

	return nil
}
