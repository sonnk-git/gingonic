package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"gingonic/graph/generated"
	"gingonic/middlewares"
	"gingonic/models"
	"github.com/gin-gonic/gin"

	//"gingonic/models"
)

// NoOp is the resolver for the NoOp field.
func (r *mutationResolver) NoOp(ctx context.Context) (*bool, error) {
	panic(fmt.Errorf("not implemented: NoOp - NoOp"))
}

// NoOp is the resolver for the NoOp field.
func (r *queryResolver) NoOp(ctx context.Context) (*bool, error) {
	panic(fmt.Errorf("not implemented: NoOp - NoOp"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

func GetUserFromContext(ctx context.Context) (models.User, error) {
	ginContext := ctx.Value("GinContextKey").(*gin.Context)
	token := ginContext.Request.Header.Get("Authentication")
	return middlewares.JwtTokenCheckInGraphql(token)
}
