package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetInfo(ctx *gin.Context) {
	id, _ := ctx.Get("id")
	email, _ := ctx.Get("email")
	username, _ := ctx.Get("username")
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"status":   true,
		"id":       id,
		"email":    email,
		"username": username,
	})
}
