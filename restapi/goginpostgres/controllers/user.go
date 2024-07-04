package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/iamSuva/postgres_restapi/config"
	"github.com/iamSuva/postgres_restapi/models"
)


func Getuser(c *gin.Context){
	id:=c.Param("id")
	var user models.User
	config.DB.Where("id= ?",id).First(&user) //find user by id
	c.JSON(200, gin.H{
		"message":"User fetched successfully",
        "data":user,
	})

}
func Getallusercontroller(c *gin.Context){
	// c.String(200,"Hello world")
	users:=[]models.User{}
	config.DB.Find(&users)
	c.JSON(200,users)

}

func Createuser(c*gin.Context){
   var user models.User
   c.BindJSON(&user)
   config.DB.Create(&user)
   c.JSON(200, gin.H{"message": "User created successfully","data":user})
}
func Updateuser(c*gin.Context){
   var user models.User
   config.DB.Where("id= ?",c.Param("id")).First(&user)
   c.BindJSON(&user)
   config.DB.Save(&user)
   c.JSON(200, gin.H{"message": "User updated successfully","data":user})
}


func Deleteuser(c* gin.Context){
	var user models.User
	id := c.Param("id")
    config.DB.Where("id=?",id).Delete( &user)
	c.JSON(200, gin.H{"message": "User deleted successfully","data":user})

}

