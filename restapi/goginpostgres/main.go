package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iamSuva/postgres_restapi/config"
	"github.com/iamSuva/postgres_restapi/routes"
)

func main() {
	router := gin.Default()
	routes.UserRoutes(router)
     //connect to database
	 config.Connect()
	//run it
	router.Run("localhost:8080")
}