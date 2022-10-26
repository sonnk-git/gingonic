package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Route(r *gin.RouterGroup) *gin.RouterGroup {
	r.GET("/info", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, struct{}{})
	})
	return r
}
