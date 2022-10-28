package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const PATH = "user"

func Route(r *gin.RouterGroup) *gin.RouterGroup {
	user := r.Group(PATH)

	user.GET("/info", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, struct{}{})
	})
	return r
}
