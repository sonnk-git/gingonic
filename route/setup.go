package route

import (
	"gingonic/db"
	"gingonic/middlewares"
	"gingonic/models"
	"gingonic/route/api"
	"gingonic/route/graph"
	"gingonic/route/web"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func SetupRouter() *gin.Engine {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Disable Console Color
	// gin.DisableConsoleColor()
	gin.SetMode(os.Getenv("MODE"))

	// use New instead of Default: https://github.com/toorop/gin-logrus/issues/2
	r := gin.New()
	r.Use(gin.Recovery())
	r = middlewares.SetUpLogger(r)
	r.Use(middlewares.CORSMiddleware())

	err = models.AutoMigrate(db.InitORM())
	if err != nil {
		log.Fatal("Error migrate DB")
	}

	r.HTMLRender = ginview.Default()
	api.RegisterAPI(r)
	web.RegisterWeb(r)
	graph.RegisterGraphQL(r)

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	authorized.POST("admin", func(c *gin.Context) {
		//user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}
