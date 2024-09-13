package main

import "log"

// CheckError logs the error but does not terminate the program
func CheckError(err error) {
	if err != nil {
		log.Printf("An error occurred: %v", err)
	}
}
