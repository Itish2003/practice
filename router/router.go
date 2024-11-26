package router

import (
	"io"
	"log"
	"obsidian/practice/service"
	"os"

	"github.com/gin-gonic/gin"
)

func NewRouter() (*gin.Engine, error) {
	r := gin.New()
	r.Use(gin.Logger())
	r.Static("/template", "./template")

	file, err := os.Create("logFile.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	gin.DefaultWriter = io.MultiWriter(file)

	r.GET("/", service.Redirectv1)

	route := r.Group("v1")
	{
		route.GET("/", service.Home)

		route.POST("/addTask", service.AddTask)

		route.GET("/getTask", service.GetAllTask)

		route.GET("/getTask/:name", service.GetTaskByName)

		route.PATCH("/updateTask/:name", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Successfully working (updateTask/:name)",
			})
		})

		route.POST("/deleteTask/:name", service.DeleteTaskByName)

		route.POST("/deleteTask", service.DeleteAllTask)
	}

	return r, nil

}
