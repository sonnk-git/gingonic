package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	model "gingonic/graph"
)

// CreateCard is the resolver for the createCard field.
func (r *mutationResolver) CreateCard(ctx context.Context, input model.NewCardInput) (*model.Card, error) {
	panic(fmt.Errorf("not implemented: CreateCard - createCard"))
}

// EditCard is the resolver for the editCard field.
func (r *mutationResolver) EditCard(ctx context.Context, input model.CardInput) (*model.Card, error) {
	panic(fmt.Errorf("not implemented: EditCard - editCard"))
}

// DeleteCard is the resolver for the deleteCard field.
func (r *mutationResolver) DeleteCard(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteCard - deleteCard"))
}

// GetCards is the resolver for the getCards field.
func (r *queryResolver) GetCards(ctx context.Context, courseID *string) ([]*model.Card, error) {
	panic(fmt.Errorf("not implemented: GetCards - getCards"))
}

// GetCard is the resolver for the getCard field.
func (r *queryResolver) GetCard(ctx context.Context, id string) (*model.Card, error) {
	panic(fmt.Errorf("not implemented: GetCard - getCard"))
}
