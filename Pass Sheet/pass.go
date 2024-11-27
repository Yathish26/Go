package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Recursive function to generate codes
func generateCodes(prefix, characters string, length int, codes *[]string) {
	if len(prefix) == length {
		*codes = append(*codes, prefix)
		return
	}

	for _, ch := range characters {
		generateCodes(prefix+string(ch), characters, length, codes)
	}
}

// Save generated codes to a file
func saveToFile(codes []string, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, code := range codes {
		_, _ = writer.WriteString(code + "\n")
	}
	writer.Flush()
	fmt.Printf("Generated codes saved to %s\n", filename)
	return nil
}

func main() {
	var length int
	var includeTextChoice, includeNumbersChoice string
	var includeText, includeNumbers bool

	// Get user input
	fmt.Print("Enter the length of the code: ")
	_, err := fmt.Scan(&length)
	if err != nil || length <= 0 {
		fmt.Println("Invalid length! Please enter a positive integer.")
		return
	}

	fmt.Print("Include text (y/n): ")
	_, _ = fmt.Scan(&includeTextChoice)
	includeText = strings.ToLower(includeTextChoice) == "y"

	fmt.Print("Include numbers (y/n): ")
	_, _ = fmt.Scan(&includeNumbersChoice)
	includeNumbers = strings.ToLower(includeNumbersChoice) == "y"

	// Create character set based on user input
	var characters string
	if includeText {
		characters += "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if includeNumbers {
		characters += "0123456789"
	}

	if len(characters) == 0 {
		fmt.Println("No characters selected! Please include text or numbers.")
		return
	}

	// Generate codes
	var codes []string
	generateCodes("", characters, length, &codes)

	// Save codes to file
	err = saveToFile(codes, "codes.txt")
	if err != nil {
		fmt.Printf("Error saving to file: %v\n", err)
	}
}
