package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/iamSuva/api_build/services"
)

type NotesControllers struct{
	noteservice services.Noteservice
}

func (n *NotesControllers) InitNoteControllerRoutes(router *gin.Engine){
	notes:=router.Group("/notes")
	notes.GET("/",n.GetallNotes())
	notes.POST("/",n.CreateNote())

}

func (n *NotesControllers) GetallNotes() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{
			// "message":"Get all notes",
			"message":n.noteservice.GetNoteserive(),
		})
	}
}
func (n *NotesControllers) CreateNote() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{
			// "message":"note is created",
			"message":n.noteservice.CreateNoteService(),
		})
	}
}