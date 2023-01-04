package v1

import (
	"gingonic/controllers/api"
	"gingonic/middlewares"
	"gingonic/route/api/v1/auth"
	"gingonic/route/api/v1/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

const V1 = "v1"

func Register(r *gin.RouterGroup) {
	/* ---------------------------  Public routes  --------------------------- */
	public := r.Group(V1)
	public.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, struct{}{})
	})
	public.POST("/send-notification", api.SendNotification)
	auth.Route(public)
	/* ---------------------------  Private routes  --------------------------- */
	private := r.Group(V1)
	private.Use(middlewares.JwtTokenCheck)
	// User
	user.Route(private)
}
