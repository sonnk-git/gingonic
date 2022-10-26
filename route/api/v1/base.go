package v1

import (
	"gingonic/route/api/v1/auth"
	"gingonic/route/api/v1/user"
	"github.com/gin-gonic/gin"
)

const V1 string = "v1"

func Register(r *gin.RouterGroup) *gin.RouterGroup {
	// Init
	v1 := r.Group(V1)

	// Authenticate
	v1 = auth.Route(v1)

	// User
	v1 = user.Route(v1)

	return v1
}
