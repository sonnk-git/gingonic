package user

import (
	"gingonic/controllers"
	"github.com/gin-gonic/gin"
)

const PATH = "user"

func Route(r *gin.RouterGroup) *gin.RouterGroup {
	user := r.Group(PATH)

	user.GET("/info", controllers.GetInfo)
	user.POST("/subscribe-notification", controllers.SubscribeNotification)
	user.POST("/set-subscribe", controllers.SetSubscribe)

	return r
}
