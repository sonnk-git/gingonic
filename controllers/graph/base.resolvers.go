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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func GetUserFromContext(ctx context.Context) (models.User, error) {
	ginContext := ctx.Value(gin.ContextKey).(*gin.Context)
	token := ginContext.Request.Header.Get("Authorization")
	return middlewares.JwtTokenCheckInGraphql(token)
}
