package todo

import (
	"os"
	"path/filepath"
	"testing"
)

func TestTodoList_SaveAndLoad(t *testing.T) {
	tmpFile := filepath.Join(t.TempDir(), "test_todo.json")

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
	tmpFile := filepath.Join(t.TempDir(), "invalid.json")
	err := os.WriteFile(tmpFile, []byte("{ invalid json "), 0644)
	if err != nil {
		t.Fatal(err)
	}

	list := &TodoList{}
	err = list.LoadFromFile(tmpFile)
	if err == nil {
		t.Error("Expected an error when loading invalid JSON, but got nil")
	}
}

func TestSaveToFile_PermissionError(t *testing.T) {
	// Attempting to save to a path that is a directory instead of a file
	// will trigger an os.WriteFile error on most systems.
	dir := filepath.Join(t.TempDir(), "testdir")
	if err := os.Mkdir(dir, 0755); err != nil {
		t.Fatal(err)
	}

	list := &TodoList{}
	err := list.SaveToFile(dir) // dir is a folder, writing fails
	if err == nil {
		t.Error("Expected error when saving to a directory path, but got nil")
	}
}

func TestTagList_SaveAndLoad(t *testing.T) {
	t.Run("Save and load tags", func(t *testing.T) {
		filename := filepath.Join(t.TempDir(), "tags.json")

		tags := &TagList{}
		tags.CreateTag("work", "blue")
		tags.CreateTag("personal", "green")

		if err := tags.SaveToFile(filename); err != nil {
			t.Fatalf("Expected no error saving tags, got %v", err)
		}

		loaded := &TagList{}

		if err := loaded.LoadFromFile(filename); err != nil {
			t.Fatalf("Expected no error loading tags, got %v", err)
		}

		if len(loaded.Tags) != 2 {
			t.Fatalf("Expected 2 tags, got %d", len(loaded.Tags))
		}

		if loaded.Tags[0].Name != "work" {
			t.Errorf("Expected first tag to be work, got %q", loaded.Tags[0].Name)
		}

		if loaded.Tags[0].Colour != "\033[38;5;75m" {
			t.Errorf(
				"Expected work tag colour to be blue, got %q",
				loaded.Tags[0].Colour,
			)
		}

		if loaded.Tags[1].Name != "personal" {
			t.Errorf("Expected second tag to be personal, got %q", loaded.Tags[1].Name)
		}
	})
}

func TestTagList_LoadFromFile_NotFound(t *testing.T) {
	tags := &TagList{}

	err := tags.LoadFromFile(
		filepath.Join(t.TempDir(), "does-not-exist.json"),
	)

	if err != nil {
		t.Fatalf("Expected no error for missing file, got %v", err)
	}

	if len(tags.Tags) != 0 {
		t.Errorf("Expected empty tag list, got %d tags", len(tags.Tags))
	}
}

func TestTagList_LoadFromFile_InvalidJSON(t *testing.T) {
	filename := filepath.Join(t.TempDir(), "invalid.json")

	if err := os.WriteFile(filename, []byte("{invalid json"), 0644); err != nil {
		t.Fatalf("Failed to create invalid JSON file: %v", err)
	}

	tags := &TagList{}

	err := tags.LoadFromFile(filename)

	if err == nil {
		t.Error("Expected error for invalid JSON, got nil")
	}
}

func TestTagList_SaveToFile_PermissionError(t *testing.T) {
	tags := &TagList{}
	tags.CreateTag("work", "blue")

	dir := t.TempDir()

	err := tags.SaveToFile(dir)

	if err == nil {
		t.Error("Expected error when saving to a directory, got nil")
	}
}
