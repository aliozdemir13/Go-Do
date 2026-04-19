package todo

import (
	"os"
	"testing"
)

// Assuming a structure like this based on your methods
type Item struct {
	Task string `json:"task"`
	Done bool   `json:"done"`
}

func TestTodoList_SaveAndLoad(t *testing.T) {
	tmpFile := "test_todo.json"
	// Cleanup after test
	defer os.Remove(tmpFile)

	list := &TodoList{
		Tasks: []Task{
			{Title: "Learn Go", IsDone: false},
			{Title: "Write Tests", IsDone: true},
		},
	}

	// 1. Test SaveToFile (Success)
	err := list.SaveToFile(tmpFile)
	if err != nil {
		t.Fatalf("Expected no error saving file, got %v", err)
	}

	// 2. Test LoadFromFile (Success)
	newList := &TodoList{}
	err = newList.LoadFromFile(tmpFile)
	if err != nil {
		t.Fatalf("Expected no error loading file, got %v", err)
	}

	if len(newList.Tasks) != 2 {
		t.Errorf("Expected 2 items, got %d", len(newList.Tasks))
	}

	if newList.Tasks[0].Title != "Learn Go" {
		t.Errorf("Data mismatch: expected 'Learn Go', got '%s'", newList.Tasks[0].Title)
	}
}

func TestLoadFromFile_NotFound(t *testing.T) {
	list := &TodoList{}
	// Test loading a file that doesn't exist
	err := list.LoadFromFile("non_existent_file.json")

	if err != nil {
		t.Errorf("Expected no error for non-existent file (should print 'Starting fresh!'), got %v", err)
	}
}

func TestLoadFromFile_InvalidJSON(t *testing.T) {
	tmpFile := "invalid.json"
	err := os.WriteFile(tmpFile, []byte("{ invalid json "), 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile)

	list := &TodoList{}
	err = list.LoadFromFile(tmpFile)
	if err == nil {
		t.Error("Expected an error when loading invalid JSON, but got nil")
	}
}

func TestSaveToFile_PermissionError(t *testing.T) {
	// Attempting to save to a path that is a directory instead of a file
	// will trigger an os.WriteFile error on most systems.
	err := os.Mkdir("testdir", 0755)
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll("testdir")

	list := &TodoList{}
	err = list.SaveToFile("testdir") // "testdir" is a folder, writing fails
	if err == nil {
		t.Error("Expected error when saving to a directory path, but got nil")
	}
}
