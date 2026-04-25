// Package todo provides functionality for the app
package todo

import (
	"errors"
	"fmt"
	"io"
	"time"
)

// TaskID tracks the index of the last task id
type TaskID int

// Task struct is responsible for details of the todo items
type Task struct {
	ID        TaskID    `json:"id"`
	Title     string    `json:"title"`
	IsDone    bool      `json:"is_done"`
	CreatedAt time.Time `json:"createdAt"`
}

// TodoList Struct (Our "Manager")
type TodoList struct { // nolint:revive
	Tasks  []Task `json:"tasks"`
	LastID int    `json:"last_id"`
}

// Add Method to Add a Task (Pointer Receiver because of modifying the list)
func (l *TodoList) Add(title string) {
	l.LastID++
	newTask := Task{
		ID:        TaskID(l.LastID),
		Title:     title,
		IsDone:    false,
		CreatedAt: time.Now(),
	}
	l.Tasks = append(l.Tasks, newTask)
}

// Display Method to List Tasks (Value Receiver because of only reading)
/* Once a struct becomes a "Manager" (like TodoList),
we almost always use Pointer Receivers (*TodoList) for everything, even if it's just reading.
It's more efficient and keeps the method set consistent. */
func (l *TodoList) Display(w io.Writer, isDone bool) {
	counter := 0
	for _, t := range l.Tasks {
		status := " "
		if t.IsDone {
			status = "X"
		}
		if t.IsDone != isDone {
			continue
		}
		counter++
		_, _ = fmt.Fprintf(w, "[%s] ID: %d | %s (Added: %v)\n", status, t.ID, t.Title, t.CreatedAt.Format("15:04:05"))

	}
	if len(l.Tasks) == 0 || (!isDone && counter == 0) {
		_, _ = fmt.Fprintln(w, "No tasks yet! Go grab a coffee.")
		return
	} else if isDone && counter == 0 {
		_, _ = fmt.Fprintln(w, "Nothing to see here! Almost too clean, lets close some tasks to fix that ;)")
		return
	}
}

// Complete function mark task as complete
func (l *TodoList) Complete(id TaskID) error {
	for i := range l.Tasks {
		if l.Tasks[i].ID == id {
			l.Tasks[i].IsDone = true
			return nil // Success! Exit early.
		}
	}
	// If we finish the loop without returning, the Id wasn't found
	return errors.New("task not found")
}

// Delete removes the task from the list
func (l *TodoList) Delete(id TaskID) error {
	for i := range l.Tasks {
		if l.Tasks[i].ID == id {
			l.Tasks = append(l.Tasks[:i], l.Tasks[i+1:]...)
			//l.Tasks[len(l.Tasks)-1] = Task{} // clean the memory from deleted items
			return nil // Success! Exit early.
		}
	}
	// If we finish the loop without returning, the Id wasn't found
	return errors.New("task not found")
}

// GetStats calculates the todo stats
func (l *TodoList) GetStats() (total int, completed int) {
	total = len(l.Tasks)
	for _, t := range l.Tasks {
		if t.IsDone {
			completed++
		}
	}
	return
}
