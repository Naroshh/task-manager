package routes

import (
	"task-manager/handlers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes initializes routes for the task manager
func SetupRoutes(r *gin.Engine) {
	r.GET("/", handlers.ListTasks)                       // Homepage for task list
	r.POST("/add", handlers.AddTask)                     // Form submission to add a new task
	r.POST("/toggle/:id", handlers.ToggleTaskCompletion) // Toggle task completion status
	r.POST("/delete/:id", handlers.DeleteTask)           // Delete a task
}
