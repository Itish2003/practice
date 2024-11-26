package service

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Base code for home
func Home(c *gin.Context) {
	tmpl, err := template.ParseFiles("template/home.html")
	if err != nil {
		log.Fatal("Couldn't parse the home HTML page:", err)
	}

	// Render the template with the current tasks
	data := gin.H{
		"Tasks": task, // Pass the dynamic task list to the template
	}
	if err := tmpl.Execute(c.Writer, data); err != nil {
		log.Println("Error executing template:", err)
	}
}

var task []string

func AddTask(c *gin.Context) {
	// Retrieve the task from the form data
	newTask := c.PostForm("task")
	if newTask == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task cannot be empty"})
		return
	}
	task = append(task, newTask)
	c.Redirect(http.StatusSeeOther, "/v1/")
}

func GetAllTask(c *gin.Context) {
	if len(task) == 0 {
		c.JSON(200, gin.H{"message": "List is empty"})
		return
	}
	c.JSON(200, gin.H{"Tasks": task})
}

func GetTaskByName(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task name is required"})
		return
	}

	for _, myTask := range task {
		if strings.EqualFold(myTask, name) {
			c.JSON(200, gin.H{"Task:": myTask})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

func DeleteAllTask(c *gin.Context) {
	task = []string{}
	// c.JSON(200, gin.H{"Tasks": task})
	c.Redirect(http.StatusSeeOther, "/v1/")
}

func UpdateTask(c *gin.Context) {

}

func DeleteTaskByName(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No such Task"})
		return
	}

	index := -1
	for i, myTask := range task {
		if strings.EqualFold(myTask, name) {
			index = i
			break
		}
	}

	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	// Remove the task from the slice
	task = append(task[:index], task[index+1:]...)
	// c.JSON(200, gin.H{"Tasks": task})
	c.Redirect(http.StatusSeeOther, "/v1/")
}
