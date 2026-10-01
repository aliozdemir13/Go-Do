package todo

import (
	"io"
	"strings"
	"testing"
)

func TestTodoList_Add(t *testing.T) {
	tags := &TagList{}
	tags.CreateTag("work", "blue")

	l := &TodoList{}

	err := l.Add("Test Task 1", "work", tags)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	err = l.Add("Test Task 2", "", tags)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(l.Tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(l.Tasks))
	}

	if l.LastID != 2 {
		t.Errorf("Expected LastID to be 2, got %d", l.LastID)
	}

	if l.Tasks[0].Title != "Test Task 1" {
		t.Errorf(
			"Expected first task title to be 'Test Task 1', got %s",
			l.Tasks[0].Title,
		)
	}

	if l.Tasks[0].Tag.Name != "work" {
		t.Errorf(
			"Expected first task tag to be 'work', got %s",
			l.Tasks[0].Tag.Name,
		)
	}

	if l.Tasks[0].Tag.Colour != "\033[38;5;75m" {
		t.Errorf(
			"Expected work tag colour to be blue, got %q",
			l.Tasks[0].Tag.Colour,
		)
	}

	if l.Tasks[1].Tag.Name != "" {
		t.Errorf(
			"Expected second task to have no tag, got %q",
			l.Tasks[1].Tag.Name,
		)
	}
}

func TestTodoList_Add_InvalidTag(t *testing.T) {
	tags := &TagList{}
	tags.CreateTag("work", "blue")

	l := &TodoList{}

	err := l.Add("Test Task", "does-not-exist", tags)
	if err == nil {
		t.Error("Expected error for non-existent tag, got nil")
	}

	if len(l.Tasks) != 0 {
		t.Errorf("Expected task not to be added, got %d tasks", len(l.Tasks))
	}
}

func TestTodoList_Complete(t *testing.T) {
	tags := &TagList{}

	l := &TodoList{}

	err := l.Add("Task to complete", "", tags)
	if err != nil {
		t.Fatalf("Expected no error adding task, got %v", err)
	}

	id := TaskID(l.LastID)

	// Success case
	err = l.Complete(id)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if !l.Tasks[0].IsDone {
		t.Error("Expected task to be marked as done")
	}

	// Failure case
	err = l.Complete(999)
	if err == nil {
		t.Error("Expected error for non-existent ID, got nil")
	}
}

func TestTodoList_Delete(t *testing.T) {
	tags := &TagList{}

	l := &TodoList{}

	if err := l.Add("Task 1", "", tags); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if err := l.Add("Task 2", "", tags); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if err := l.Add("Task 3", "", tags); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Delete middle task (Task 2, ID 2)
	err := l.Delete(2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(l.Tasks) != 2 {
		t.Errorf("Expected 2 tasks remaining, got %d", len(l.Tasks))
	}

	if l.Tasks[1].Title != "Task 3" {
		t.Error("Task 3 should have shifted up to index 1")
	}

	// Delete non-existent
	err = l.Delete(999)
	if err == nil {
		t.Error("Expected error for non-existent ID, got nil")
	}
}

func TestTodoList_GetStats(t *testing.T) {
	tags := &TagList{}

	l := &TodoList{}

	if err := l.Add("T1", "", tags); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if err := l.Add("T2", "", tags); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if err := l.Add("T3", "", tags); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if err := l.Complete(1); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if err := l.Complete(2); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	total, completed := l.GetStats()

	if total != 3 {
		t.Errorf("Expected total 3, got %d", total)
	}

	if completed != 2 {
		t.Errorf("Expected completed 2, got %d", completed)
	}
}

func TestTodoList_Display(t *testing.T) {
	t.Run("Empty List", func(t *testing.T) {
		l := &TodoList{}

		l.Display(io.Discard, false, "")
	})

	t.Run("No Completed Tasks Message", func(t *testing.T) {
		tags := &TagList{}
		l := &TodoList{}

		if err := l.Add("Pending", "", tags); err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		l.Display(io.Discard, true, "")
	})

	t.Run("No Pending Tasks Message", func(t *testing.T) {
		tags := &TagList{}
		l := &TodoList{}

		if err := l.Add("Done", "", tags); err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if err := l.Complete(1); err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		l.Display(io.Discard, false, "")
	})

	t.Run("Full Display", func(t *testing.T) {
		tags := &TagList{}
		l := &TodoList{}

		if err := l.Add("Task A", "", tags); err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if err := l.Add("Task B", "", tags); err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if err := l.Complete(1); err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		l.Display(io.Discard, true, "")
		l.Display(io.Discard, false, "")
	})
}

func TestStyleTextWithTagName(t *testing.T) {
	t.Run("Known tag color", func(t *testing.T) {
		got := StyleTextWithTagName("blue", "work")
		want := "\033[38;5;75mwork" + ColorReset

		if got != want {
			t.Errorf("Expected %q, got %q", want, got)
		}
	})

	t.Run("Unknown tag color", func(t *testing.T) {
		got := StyleTextWithTagName("does-not-exist", "work")
		want := "work"

		if got != want {
			t.Errorf("Expected %q, got %q", want, got)
		}
	})
}

func TestPrintTags(t *testing.T) {
	t.Run("Print tags", func(t *testing.T) {
		tags := &TagList{}

		tags.CreateTag("work", "blue")
		tags.CreateTag("personal", "green")

		var output strings.Builder

		err := PrintTags(&output, tags)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		got := output.String()

		if !strings.Contains(got, "Tags:") {
			t.Error("Expected output to contain Tags:")
		}

		if !strings.Contains(got, "work") {
			t.Error("Expected output to contain work tag")
		}

		if !strings.Contains(got, "personal") {
			t.Error("Expected output to contain personal tag")
		}
	})

	t.Run("Empty tags", func(t *testing.T) {
		tags := &TagList{}

		var output strings.Builder

		err := PrintTags(&output, tags)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if !strings.Contains(output.String(), "Tags:") {
			t.Error("Expected output to contain Tags:")
		}
	})
}
