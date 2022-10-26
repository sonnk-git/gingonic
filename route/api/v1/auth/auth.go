package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Route(r *gin.RouterGroup) *gin.RouterGroup {

	r.POST("/login", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, struct{}{})
	})

	r.POST("/register", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, struct{}{})
	})

	r.POST("/logout", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, struct{}{})
	})

	r.POST("/forgot-password", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, struct{}{})
	})

	return r
}
