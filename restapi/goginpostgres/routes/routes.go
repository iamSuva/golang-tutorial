package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/iamSuva/postgres_restapi/controllers"
)

func UserRoutes(router *gin.Engine) {
    //  router.GET("/get",func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"message": "Get all users",
	// 		"status": 200,
	// 	})
	//  })
	router.GET("/get",controllers.Getallusercontroller)
	router.GET("/get/:id",controllers.Getuser)
	router.POST("/post",controllers.Createuser)
	router.PUT("/put/:id",controllers.Updateuser)
	router.DELETE("/delete/:id",controllers.Deleteuser)
}