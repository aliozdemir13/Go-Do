// Package todo provides functionality for the app
package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

// SaveToFile saves the tasks to a file
func (l *TodoList) SaveToFile(filename string) error {
	// Convert the struct into a JSON byte slice
	// MarshalIndent makes the file "pretty"
	data, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return err
	}

	// Write the bytes to a file
	// 0644 is a Linux permission code (standard read/write) - !!requires for saving, otherwise permission issue may arise
	return os.WriteFile(filename, data, 0644)
}

// LoadFromFile fetches the tasks from file
func (l *TodoList) LoadFromFile(filename string) error {
	// Read the file
	data, err := os.ReadFile(filename)
	if os.IsNotExist(err) {
		fmt.Printf("Starting fresh!")
		return nil
	}

	return json.Unmarshal(data, l)
}
