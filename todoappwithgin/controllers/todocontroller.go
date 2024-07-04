package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamSuva/todoapp/config"
	"github.com/iamSuva/todoapp/models"
)

func GetTodos(ctx *gin.Context){
	var todos []models.Todo
	config.DB.Find(&todos) //find all todos
	ctx.JSON(http.StatusOK,gin.H{
		"message":"Get all todos",
		"data":todos,
        "status": 200,
	})
}
func GetTodoById(ctx *gin.Context){
	var todo models.Todo
  
	id:=ctx.Param("id")
	if err:=config.DB.Where("id=?",id).First(&todo).Error;err!=nil {
		ctx.JSON(http.StatusNotFound, gin.H{
            "message": "Todo not found",
            "status":  404,
        })
        return  
    }
    //todo not found, return 404 error.
	ctx.JSON(http.StatusOK,gin.H{
		"message":"Get singal todo",
		"status": 200,
		"data":todo,
	})
}

func CreateTodo(ctx *gin.Context){
	var newTodo models.Todo
	if err := ctx.ShouldBindJSON(&newTodo); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "message": "Invalid input",
            "status":  400,
        })
        return
    }
	 // Save the new todo
	 if err := config.DB.Create(&newTodo).Error; err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "message": "Failed to create todo",
            "status":  500,
        })
        return
    }
	ctx.JSON(http.StatusOK,gin.H{
        "message":"Create todo",
		"data":newTodo,
        "status": 200,
    })
}
func UpdateTodo(ctx *gin.Context){
	var updatedTodo models.Todo
	id:=ctx.Param("id")
	if err:=config.DB.Where("id=?",id).First(&updatedTodo).Error;err!=nil{
		ctx.JSON(http.StatusNotFound, gin.H{
            "message": "Todo not found",
            "status":  404,
        })
        return  
		//todo not found, return 404 error.
		//  404 means not found in this case.
		//   400 means bad request. 500 means server error. 
		// 200 means OK. 201 means created. 202 means accepted.
		//  204 means no content. 301 means moved permanently. 
	}
	ctx.BindJSON(&updatedTodo)
	config.DB.Save(&updatedTodo) //update todo by id

	ctx.JSON(http.StatusOK,gin.H{
        "message":"Update todo",
		"data":updatedTodo,
        "status": 200,
    })
}
func DeleteTodo(ctx *gin.Context){
	var deletedTodo models.Todo
	id:=ctx.Param("id")
	if err:=config.DB.Where("id=?",id).First(&deletedTodo).Error;err!=nil{
		ctx.JSON(http.StatusNotFound, gin.H{
            "message": "Todo not found",
            "status":  404,
        })
        return
	} //find todo by id
	config.DB.Delete(&deletedTodo) //delete todo by id
	ctx.JSON(http.StatusOK,gin.H{
        "message":"Delete todo",
		"data":deletedTodo,
        "status": 200,
    })
}