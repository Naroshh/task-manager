// main.go
package main

import (
	"log"
	"net/http"
	"task-manager/handlers"
)

func main() {
	// Route handlers
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/tasks", handlers.TasksHandler)
	http.HandleFunc("/tasks/add", handlers.AddTaskHandler)
	http.HandleFunc("/tasks/delete", handlers.DeleteTaskHandler)
	http.HandleFunc("/tasks/toggle", handlers.ToggleTaskHandler) // Correct handler name

	// Start server
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
