package api

import (
	v1 "gingonic/route/api/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterAPI(r *gin.Engine) {
	apiGroupRouter := r.Group("/api")
	//apiGroupRouter.Use(middlewares.JwtTokenCheck)

	apiGroupRouter.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, struct{}{})
	})

	v1.Register(apiGroupRouter)
}
