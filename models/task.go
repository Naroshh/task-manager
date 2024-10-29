// models/task.go
package models

import "sync"

// Task represents a single task with its details.
type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"dueDate"`
	Priority    string `json:"priority"` // e.g., "Low", "Medium", "High"
	Status      string `json:"status"`   // "Incomplete" or "Complete"
}

var (
	tasks   = []Task{}
	nextID  = 1
	taskMux sync.Mutex
)

// GetTasks returns all tasks.
func GetTasks() []Task {
	taskMux.Lock()
	defer taskMux.Unlock()
	return tasks
}

// AddTask adds a new task with the provided details.
func AddTask(title, description, dueDate, priority string) Task {
	taskMux.Lock()
	defer taskMux.Unlock()

	task := Task{
		ID:          nextID,
		Title:       title,
		Description: description,
		DueDate:     dueDate,
		Priority:    priority,
		Status:      "Incomplete", // New tasks are incomplete by default
	}
	nextID++
	tasks = append(tasks, task)
	return task
}

// DeleteTask deletes a task by ID.
func DeleteTask(id int) bool {
	taskMux.Lock()
	defer taskMux.Unlock()

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...) // Remove task from slice
			return true
		}
	}
	return false
}

// ToggleTask toggles the completion status of a task.
func ToggleTask(id int) bool {
	taskMux.Lock()
	defer taskMux.Unlock()

	for i, task := range tasks {
		if task.ID == id {
			if task.Status == "Incomplete" {
				tasks[i].Status = "Complete" // Mark as complete
			} else {
				tasks[i].Status = "Incomplete" // Mark as incomplete
			}
			return true
		}
	}
	return false
}
