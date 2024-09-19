package utils

import (
	"log"
	"os"
)

// CheckError logs the error but does not terminate the program
func CheckError(err error) {
	if err != nil {
		log.Printf("An error occurred: %v", err)
	}
}

// FatalError logs the error and terminates the program
func FatalError(err error) {
	if err != nil {
		log.Fatalf("Fatal error: %v", err)
		os.Exit(1) // Ensure the program exits after logging the fatal error
	}
}

// InfoLog logs informational messages for debugging or tracing purposes
func InfoLog(message string) {
	log.Printf("Info: %s", message)
}

// FileExists checks if a file exists at the given path
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

// CreateFile creates a new file and returns the file object
func CreateFile(filename string) (*os.File, error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	InfoLog("File created: " + filename)
	return file, nil
}

// ReadFile reads the contents of a file into a byte slice
func ReadFile(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// WriteToFile writes the given data to a file
func WriteToFile(filename string, data []byte) error {
	err := os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	InfoLog("Data written to file: " + filename)
	return nil
}
