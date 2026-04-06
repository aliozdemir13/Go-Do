package todo

import (
	"errors"
	"fmt"
	"time"
)

type TaskId int

type Task struct {
	Id        TaskId    `json:"id"`
	Title     string    `json:"title"`
	IsDone    bool      `json:"is_done"`
	CreatedAt time.Time `json:"createdAt"`
}

// The TodoList Struct (Our "Manager")
type TodoList struct {
	Tasks  []Task `json:"tasks"`
	LastId int    `json:"last_id"`
}

// Method to Add a Task (Pointer Receiver because of modifying the list)
func (l *TodoList) Add(title string) {
	l.LastId++
	newTask := Task{
		Id:        TaskId(l.LastId),
		Title:     title,
		IsDone:    false,
		CreatedAt: time.Now(),
	}
	l.Tasks = append(l.Tasks, newTask)
}

// Method to List Tasks (Value Receiver because of only reading)
/* Once a struct becomes a "Manager" (like TodoList),
we almost always use Pointer Receivers (*TodoList) for everything, even if it's just reading.
It's more efficient and keeps the method set consistent. */
func (l *TodoList) Display(isDone bool) {
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
		fmt.Printf("[%s] ID: %d | %s (Added: %v)\n", status, t.Id, t.Title, t.CreatedAt.Format("15:04:05"))

	}
	if len(l.Tasks) == 0 || (!isDone && counter == 0) {
		fmt.Println("No tasks yet! Go grab a coffee.")
		return
	} else if isDone && counter == 0 {
		fmt.Println("Nothing to see here! Almost too clean, lets close some tasks to fix that ;)")
		return
	}
}

func (l *TodoList) Complete(id TaskId) error {
	for i := range l.Tasks {
		if l.Tasks[i].Id == id {
			l.Tasks[i].IsDone = true
			return nil // Success! Exit early.
		}
	}
	// If we finish the loop without returning, the Id wasn't found
	return errors.New("task not found")
}

func (l *TodoList) Delete(id TaskId) error {
	for i := range l.Tasks {
		if l.Tasks[i].Id == id {
			l.Tasks = append(l.Tasks[:i], l.Tasks[i+1:]...)
			//l.Tasks[len(l.Tasks)-1] = Task{} // clean the memory from deleted items
			return nil // Success! Exit early.
		}
	}
	// If we finish the loop without returning, the Id wasn't found
	return errors.New("task not found")
}

func (l *TodoList) GetStats() (total int, completed int) {
	total = len(l.Tasks)
	for _, t := range l.Tasks {
		if t.IsDone {
			completed++
		}
	}
	return
}
