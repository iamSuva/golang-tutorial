package middleware

import "github.com/gin-gonic/gin"

func Authenticate(ctx *gin.Context){
	if !(ctx.Request.Header.Get("Token")=="auth"){
		ctx.AbortWithStatusJSON(401,gin.H{"error":"Unauthorized token is not present"})
        return
	}
	ctx.Next()
}

func AddHeader(ctx *gin.Context){
	ctx.Writer.Header().Set("name","Suvadip")  // this will add a header to the response object {name:suvadip}
	ctx.Next()
}