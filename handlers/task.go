// handlers/task.go
package handlers

import (
	"encoding/json"
	"net/http"
	"task-manager/models"
)

// TasksHandler returns all tasks as JSON.
func TasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tasks := models.GetTasks()
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		http.Error(w, "Error encoding tasks", http.StatusInternalServerError)
		return
	}
}

// AddTaskHandler adds a new task from JSON data.
func AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		DueDate     string `json:"dueDate"`
		Priority    string `json:"priority"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	task := models.AddTask(request.Title, request.Description, request.DueDate, request.Priority)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, "Error encoding task", http.StatusInternalServerError)
		return
	}
}

// DeleteTaskHandler deletes a task by ID.
func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	var request struct{ ID int }
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if models.DeleteTask(request.ID) {
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Task not found", http.StatusNotFound)
	}
}

// ToggleTaskHandler toggles the completion status of a task by ID.
func ToggleTaskHandler(w http.ResponseWriter, r *http.Request) {
	var request struct{ ID int }
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if models.ToggleTask(request.ID) {
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Task not found", http.StatusNotFound)
	}
}
