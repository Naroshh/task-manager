package models

// Task represents a single task in our task manager
type Task struct {
	ID        int    // Unique ID for the task
	Title     string // Title or description of the task
	Completed bool   // Whether the task is completed
}
