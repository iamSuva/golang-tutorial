package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iamSuva/todoapp/config"
	"github.com/iamSuva/todoapp/routes"
)

func main() {
	router := gin.Default()
	config.Connect()

	routes.TodoRoutes(router)
	router.Run("localhost:8080")
}
