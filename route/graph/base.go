package graph

import (
	"context"
	"fmt"
	"gingonic/controllers/graph"
	"gingonic/graph/generated"
	"gingonic/middlewares"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"os"
	"time"
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
		ginContext := ctx.Value(gin.ContextKey).(*gin.Context)
		token := ginContext.Request.Header.Get("Authorization")
		_, err = middlewares.JwtTokenCheckInGraphql(token)
		if err != nil {
			return nil, fmt.Errorf("%v", err.Error())
		}
		return next(ctx)
	}

	// when using NewDefaultServer, the default uses SameOrigin. So if you're running your client on a different port it won't upgrade
	// reference link https://github.com/99designs/gqlgen/issues/1328#issuecomment-742212770
	// https://github.com/99designs/gqlgen/blob/master/docs/content/recipes/authentication.md
	h := handler.New(generated.NewExecutableSchema(c))
	h.AddTransport(&transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		InitFunc: func(ctx context.Context, initPayload transport.InitPayload) (context.Context, error) {
			return webSocketInit(ctx, initPayload)
		},
	})

	h.AddTransport(transport.Options{})
	h.AddTransport(transport.GET{})
	h.AddTransport(transport.POST{})
	h.AddTransport(transport.MultipartForm{})

	h.SetQueryCache(lru.New(1000))

	h.Use(extension.Introspection{})
	h.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	return func(c *gin.Context) {
		if os.Getenv("MODE") == gin.ReleaseMode {
			graphql.GetOperationContext(c).DisableIntrospection = true
		}

		h.ServeHTTP(c.Writer, c.Request)
	}
}

func RegisterGraphQL(r *gin.Engine) {
	r.Use(ginContextToContextMiddleware())

	r.Any("/graphql", graphqlHandler())
	r.GET("/playground", playgroundHandler())
}

func ginContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), gin.ContextKey, c)
		c.Request = c.Request.WithContext(ctx)
		c.Set("token", c.Request.Header.Get("Authorization"))
		c.Next()
	}
}

func webSocketInit(ctx context.Context, initPayload transport.InitPayload) (context.Context, error) {
	// Get the token from payload
	//any := initPayload["authToken"]
	//token, ok := any.(string)
	//if !ok || token == "" {
	//	return nil, errors.New("authToken not found in transport payload")
	//}
	//
	//// Perform token verification and authentication...
	//userId := "john.doe" // e.g. userId, err := GetUserFromAuthentication(token)

	// put it in context
	ctxNew := context.WithValue(ctx, "username", 1)

	return ctxNew, nil
}
