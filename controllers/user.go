package controllers

import "github.com/gin-gonic/gin"

func GetInfo(ctx *gin.Context)  {
	token := ctx.Request.Header.Get("Authorization")
	println(token)
}
