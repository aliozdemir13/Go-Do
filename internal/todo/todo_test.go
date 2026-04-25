package todo

import (
	"io"
	"testing"
)

func TestTodoList_Add(t *testing.T) {
	l := &TodoList{}
	l.Add("Test Task 1")
	l.Add("Test Task 2")

	if len(l.Tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(l.Tasks))
	}
	if l.LastID != 2 {
		t.Errorf("Expected LastId to be 2, got %d", l.LastID)
	}
	if l.Tasks[0].Title != "Test Task 1" {
		t.Errorf("Expected first task title to be 'Test Task 1', got %s", l.Tasks[0].Title)
	}
}

func TestTodoList_Complete(t *testing.T) {
	l := &TodoList{}
	l.Add("Task to complete")
	id := TaskID(l.LastID)

	// Success case
	err := l.Complete(id)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !l.Tasks[0].IsDone {
		t.Error("Expected task to be marked as done")
	}

	// Failure case (ID doesn't exist)
	err = l.Complete(999)
	if err == nil {
		t.Error("Expected error for non-existent ID, got nil")
	}
}

func TestTodoList_Delete(t *testing.T) {
	l := &TodoList{}
	l.Add("Task 1")
	l.Add("Task 2")
	l.Add("Task 3")

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
	l := &TodoList{}
	l.Add("T1")
	l.Add("T2")
	l.Add("T3")
	l.Complete(1)
	l.Complete(2)

	total, completed := l.GetStats()
	if total != 3 {
		t.Errorf("Expected total 3, got %d", total)
	}
	if completed != 2 {
		t.Errorf("Expected completed 2, got %d", completed)
	}
}

func TestTodoList_Display(t *testing.T) {
	// We test all branches of the Display logic

	t.Run("Empty List", func(t *testing.T) {
		l := &TodoList{}
		l.Display(io.Discard, false) // Should hit "No tasks yet!"
	})

	t.Run("No Completed Tasks Message", func(t *testing.T) {
		l := &TodoList{}
		l.Add("Pending")
		l.Display(io.Discard, true) // isDone=true, but counter=0. Hits "Nothing to see here!"
	})

	t.Run("No Pending Tasks Message", func(t *testing.T) {
		l := &TodoList{}
		l.Add("Done")
		l.Complete(1)
		l.Display(io.Discard, false) // isDone=false, but counter=0. Hits "No tasks yet!"
	})

	t.Run("Full Display", func(t *testing.T) {
		l := &TodoList{}
		l.Add("Task A")
		l.Add("Task B")
		l.Complete(1)

		l.Display(io.Discard, true)  // Displays Task A
		l.Display(io.Discard, false) // Displays Task B
	})
}
