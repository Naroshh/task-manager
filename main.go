package main

import (
	"task-manager/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*") // Make sure your HTML files are in the 'templates' folder

	// Setup routes
	r.GET("/", handlers.ListTasks)                       // Homepage for task list
	r.POST("/add", handlers.AddTask)                     // Form submission to add a new task
	r.POST("/toggle/:id", handlers.ToggleTaskCompletion) // Toggle task completion status
	r.POST("/delete/:id", handlers.DeleteTask)           // Delete a task by ID

	// Start the server
	r.Run(":8080") // Listen on port 8080
}
