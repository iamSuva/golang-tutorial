package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamSuva/middleware/middleware"
)

func main() {
	router := gin.Default()

	//apply  middleware to all
	//router.Use(middleware.Authenticate)
	router.GET("/admin",GetAdmin)
    router.GET("/client",GetClient)
	//apply to paricular route
	router.GET("/auth",middleware.Authenticate,middleware.AddHeader,Getdata)
	//apply middleware to group
	user:=router.Group("/user",middleware.Authenticate)
	{
		user.GET("/getdata",Getdata)
	}
	router.Run("localhost:8080")
}
func GetAdmin(ctx *gin.Context){
	ctx.JSON(http.StatusOK,gin.H{
		"message":"Welcome admin",
	})
}
func GetClient(ctx *gin.Context){
	ctx.JSON(http.StatusOK,gin.H{
		"message":"Welcome client",
	})
}

func Getdata(ctx *gin.Context){
	ctx.JSON(http.StatusOK,gin.H{
        "message":"Welcome authenticated user",
    })
    
}