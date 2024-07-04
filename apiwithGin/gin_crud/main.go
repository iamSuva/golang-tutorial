package main

import (
	"fmt"
	// "net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamSuva/api_build/controllers"
	internal "github.com/iamSuva/api_build/internal/database"
)

func main() {
	router := gin.Default()
   db:=internal.InitDB()
   if db==nil{
	    fmt.Println("Error initializing database")
		return
   }
// 	router.GET("/api",func(ctx *gin.Context) {
// 		ctx.JSON(http.StatusOK,gin.H{
// 			"message":"Hello welcome to golang gin webframework",
// 			"status":http.StatusOK,
// 		})
// 	})
//    router.GET("/api/:id/:id2",func(ctx *gin.Context) {
// 	  id1:=ctx.Param("id")
// 	  id2:=ctx.Param("id2")
// 	  ctx.JSON(http.StatusOK,gin.H{
// 		"userid":id1,
// 		"userid2":id2,
// 	  })
//    })

//   //post request
//   router.POST("/api/post",func(ctx *gin.Context) {
//     type User struct{
// 		Username string `json:"username" binding:"required"`
// 		Password string `json:"password" binding:"required"`
// 	}
// 	var user User
// 	if err:=ctx.BindJSON(&user);err!=nil{
// 		ctx.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
//         return   
//   }
	
// 	ctx.JSON(http.StatusOK,gin.H{
// 		"message":"User created successfully",
//         "user":user,
// 		"username":user.Username,
// 	})
    

// })
// router.PUT("/api/put",func(ctx *gin.Context) {
//     type User struct{
// 		Username string `json:"username" binding:"required"`
// 		Password string `json:"password" binding:"required"`
// 	}
// 	var user User
// 	if err:=ctx.BindJSON(&user);err!=nil{
// 		ctx.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
//         return   
//   }
	
// 	ctx.JSON(http.StatusOK,gin.H{
// 		"message":"put request made successfully",
//         "user":user,
// 		"username":user.Username,
// 	})
    

// })
// router.PATCH("/api/patch",func(ctx *gin.Context) {
//     type User struct{
// 		Username string `json:"username" `
// 		Password string `json:"password" `
// 	}
// 	var user User
// 	if err:=ctx.BindJSON(&user);err!=nil{
// 		ctx.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
//         return   
//   }
	
// 	ctx.JSON(http.StatusOK,gin.H{
// 		"message":"patch request made  successfully",
//         "user":user,
// 		"username":user.Username,
// 	})
    

// })

// router.DELETE("api/delete/:id",func(ctx *gin.Context) {
// 	id:=ctx.Param("id")
//     ctx.JSON(http.StatusOK,gin.H{
//         "message":"User with id "+id+" deleted successfully",
//     })
    
//     //you can add database logic here to delete user from database
// })
  
   notescontroller:=&controllers.NotesControllers{}
   notescontroller.InitNoteControllerRoutes(router)
	fmt.Println("server is running on port 8080")
	router.Run("localhost:8080")
}
