package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	DueDate     time.Time `json:"due_date"`
}

var tasks = []Task{
	{ID: "1", Title: "Task 1", Description: "Task 1 description",DueDate:time.Now(), Status: "Pending"},
	{ID: "2", Title: "Task 2", Description: "Task 2 description",DueDate: time.Now().AddDate(0,0,2), Status: "Pending"},
}

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "welcome to go lang"})
	})
	router.GET("/tasks", getAlltasks)
	router.GET("/tasks/:id", singletask)
	router.POST("/tasks", addTask)
	router.PUT("/tasks/:id", updateTask)
	router.DELETE("/tasks/:id", deleteTask)

	fmt.Println("server running on port 4000")
	router.Run(":4000")
}
func getAlltasks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, tasks)
}

func singletask(ctx *gin.Context) {
	id := ctx.Param("id")
	for _, task := range tasks {
		if task.ID == id {
			ctx.JSON(http.StatusOK, task)
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"message": "task not found"})

}
func updateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	var updatedTask Task
	if err := ctx.BindJSON(&updatedTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

	}
	for i, task := range tasks {
		if task.ID == id {
			if updatedTask.Title != "" {
				tasks[i].Title = updatedTask.Title

			}
			if updatedTask.Description != "" {
				tasks[i].Description = updatedTask.Description
			}
			if updatedTask.Status != "" {
				tasks[i].Status = updatedTask.Status
			}
			ctx.JSON(http.StatusOK, gin.H{"message": tasks[i]})
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"message": "task not found"})
}

func deleteTask(ctx *gin.Context) {
	id := ctx.Param("id")
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			ctx.JSON(http.StatusOK, gin.H{"message": "task is deleted"})
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"message": "task not found"})
}

func addTask(ctx *gin.Context) {
	var newTask Task
	if err := ctx.ShouldBindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	newTask.DueDate = time.Now().AddDate(0, 0, 7)

	tasks = append(tasks, newTask)
	ctx.JSON(http.StatusCreated, newTask)
}
