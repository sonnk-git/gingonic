package user

import (
	"gingonic/controllers/api"
	"github.com/gin-gonic/gin"
)

const PATH = "user"

func Route(r *gin.RouterGroup) *gin.RouterGroup {
	user := r.Group(PATH)

	user.GET("/info", api.GetInfo)
	user.POST("/subscribe-notification", api.SubscribeNotification)
	user.POST("/set-subscribe", api.SetSubscribe)

	return r
}
