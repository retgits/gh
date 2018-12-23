// Package util implements utility methods
package util

import (
	"fmt"
	"os"
	"path/filepath"
)

// GetCurrentDirectory gets the directory in which the app was started and returns either
// the full directory or an error
func GetCurrentDirectory() (string, error) {
	return filepath.Abs(filepath.Dir(os.Args[0]))
}

// WriteFile writes the content, passed in as a parameter, to a file, which is also passed in as a parameter. It tries to create the file if it doesn't exist and will return errors when one occurs.
func WriteFile(filename string, content string) error {
	// Create a file on disk
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("error while creating file: %s", err.Error())
		return fmt.Errorf("error while creating file: %s", err.Error())
	}
	defer file.Close()

	// Open the file to write
	file, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Printf("error while opening file: %s", err.Error())
		return fmt.Errorf("error while opening file: %s", err.Error())
	}

	// Write the content to disk
	_, err = file.Write([]byte(content))
	if err != nil {
		fmt.Printf("error while writing Markdown to disk: %s", err.Error())
		return fmt.Errorf("error while writing Markdown to disk: %s", err.Error())
	}

	return nil
}
