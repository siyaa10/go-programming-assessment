package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Function to write a message into a file
func WriteToFile(filename string, message string) error {

	// Open or create file in append mode
	file, err := os.OpenFile(
		filename,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)

	if err != nil {
		return err
	}

	// Close file when function exits
	defer file.Close()

	// Write message with a newline
	_, err = file.WriteString(message + "\n")

	if err != nil {
		return err
	}

	return nil
}

// Interactive demonstration
func runFileDemo() {

	fmt.Println("\n===== FILE HANDLING =====")

	fmt.Scanln()
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter filename: ")
	filename, _ := reader.ReadString('\n')
	filename = strings.TrimSpace(filename)

	fmt.Print("Enter message to save: ")
	message, _ := reader.ReadString('\n')
	message = strings.TrimSpace(message)

	err := WriteToFile(filename, message)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Message saved successfully!")
	}
}
