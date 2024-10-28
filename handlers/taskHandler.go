package handlers

import (
	"net/http"
	"strconv"
	"task-manager/models"

	"github.com/gin-gonic/gin"
)

// In-memory storage for tasks
var tasks = []models.Task{}

// ListTasks renders the index page with all tasks
func ListTasks(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"tasks": tasks,
	})
}

// AddTask adds a new task based on form input
func AddTask(c *gin.Context) {
	title := c.PostForm("title")
	if title == "" {
		c.String(http.StatusBadRequest, "Task title cannot be empty")
		return
	}

	// Create and add new task to the list
	task := models.Task{
		ID:        len(tasks) + 1,
		Title:     title,
		Completed: false,
	}
	tasks = append(tasks, task)
	c.Redirect(http.StatusSeeOther, "/")
}

// ToggleTaskCompletion toggles the completed status of a task
func ToggleTaskCompletion(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid task ID")
		return
	}

	// Toggle completion status
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Completed = !task.Completed
			break
		}
	}
	c.Redirect(http.StatusSeeOther, "/")
}

// DeleteTask deletes a task by its ID
func DeleteTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid task ID")
		return
	}

	// Delete task by filtering it out
	var updatedTasks []models.Task
	for _, task := range tasks {
		if task.ID != id {
			updatedTasks = append(updatedTasks, task)
		}
	}
	tasks = updatedTasks
	c.Redirect(http.StatusSeeOther, "/")
}
