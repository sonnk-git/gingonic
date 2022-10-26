package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterWeb(r *gin.Engine) {
	r.GET("/page", func(ctx *gin.Context) {
		//render only file, must full name with extension
		ctx.HTML(http.StatusOK, "page.html", gin.H{"title": "Page file title!!"})
	})

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"user": "user", "status": "no value"})
	})
}
