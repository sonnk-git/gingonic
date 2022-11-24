package graph

import (
	"context"
	"fmt"
	"gingonic/controllers/graph"
	"gingonic/graph/generated"
	"gingonic/middlewares"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"os"
)

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/graphql")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	c := generated.Config{Resolvers: &graph.Resolver{}}
	c.Directives.Authenticated = func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
		var ginContext *gin.Context
		ginContext = ctx.Value("GinContextKey").(*gin.Context)
		token := ginContext.Request.Header.Get("Authorization")
		_, err = middlewares.JwtTokenCheckInGraphql(token)
		if err != nil {
			return nil, fmt.Errorf("%v", err.Error())
		}
		return next(ctx)
	}
	h := handler.NewDefaultServer(generated.NewExecutableSchema(c))
	return func(c *gin.Context) {
		if os.Getenv("MODE") == gin.ReleaseMode {
			graphql.GetOperationContext(c).DisableIntrospection = true
		}

		h.ServeHTTP(c.Writer, c.Request)
	}
}

func RegisterGraphQL(r *gin.Engine) {
	r.Use(ginContextToContextMiddleware())

	r.POST("/graphql", graphqlHandler())
	r.GET("/playground", playgroundHandler())
}

func ginContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "GinContextKey", c)
		c.Request = c.Request.WithContext(ctx)
		c.Set("token", c.Request.Header.Get("Authorization"))
		c.Next()
	}
}
