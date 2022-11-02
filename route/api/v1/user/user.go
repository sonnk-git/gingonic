package user

import (
	"gingonic/controllers"
	"github.com/gin-gonic/gin"
)

const PATH = "user"

func Route(r *gin.RouterGroup) *gin.RouterGroup {
	user := r.Group(PATH)

	user.GET("/info", controllers.GetInfo)

	return r
}
