package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/iamSuva/todoapp/controllers"
)

func TodoRoutes(router *gin.Engine) {
	router.GET("/todos",controllers.GetTodos)
	router.GET("/todos/:id",controllers.GetTodoById)
	router.POST("/todos", controllers.CreateTodo)
	router.PUT("/todos/:id", controllers.UpdateTodo)
	router.DELETE("/todos/:id", controllers.DeleteTodo)

}