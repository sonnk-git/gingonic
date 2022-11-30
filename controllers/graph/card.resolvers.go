package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"gingonic/db"
	model "gingonic/graph"
	OrmModels "gingonic/models"
	"strings"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

// CreateCard is the resolver for the createCard field.
func (r *mutationResolver) CreateCard(ctx context.Context, input model.NewCardInput) (*model.Card, error) {
	user, err := GetUserFromContext(ctx)
	if err != nil {
		return nil, gqlerror.Errorf("Error when get user from context")
	}

	course := &OrmModels.Course{}
	tx := db.Orm.First(&course, "id = ?", input.CourseID)
	if tx.Error != nil || tx.RowsAffected < 1 {
		return nil, gqlerror.Errorf("Error when get course in CreateCard")
	}
	if user.ID != course.UserID {
		return nil, gqlerror.Errorf("User not allow to create card to this course")
	}

	card := &OrmModels.Card{
		Terminology: *input.Terminology,
		Definition:  *input.Definition,
		CourseID:    course.ID,
	}

	tx = db.Orm.Create(&card)
	if tx.Error != nil {
		return nil, gqlerror.Errorf("Error when create card to db")
	}

	cardGQL := &model.Card{
		ID:          card.ID,
		Terminology: &card.Terminology,
		Definition:  &card.Definition,
		CourseID:    card.CourseID,
	}
	return cardGQL, nil
}

// EditCard is the resolver for the editCard field.
func (r *mutationResolver) EditCard(ctx context.Context, input model.CardInput) (*model.Card, error) {
	user, err := GetUserFromContext(ctx)
	if err != nil {
		return nil, gqlerror.Errorf("Error when get user from context")
	}

	card := OrmModels.Card{}
	tx := db.Orm.First(&card, "id = ?", input.ID)
	if tx.Error != nil {
		return nil, gqlerror.Errorf("Error when get card in EditCard")
	}

	course := OrmModels.Course{}
	tx = db.Orm.First(&course, "id = ?", card.CourseID)
	if tx.Error != nil {
		return nil, gqlerror.Errorf("Error when get card in EditCard")
	}
	if course.UserID != user.ID {
		return nil, gqlerror.Errorf("User not allow to edit card in course " + course.ID)
	}

	tx = db.Orm.Model(&card).Updates(OrmModels.Card{
		Terminology: *input.Terminology,
		Definition:  *input.Definition,
	})

	if tx.Error != nil {
		return nil, gqlerror.Errorf("Error when update card in EditCard")
	}

	cardORM := model.Card{
		ID:          card.ID,
		Terminology: &card.Terminology,
		Definition:  &card.Definition,
		CourseID:    card.CourseID,
	}

	return &cardORM, nil
}

// DeleteCard is the resolver for the deleteCard field.
func (r *mutationResolver) DeleteCard(ctx context.Context, id string) (bool, error) {
	user, err := GetUserFromContext(ctx)
	if err != nil {
		return false, gqlerror.Errorf("Error when get user from context")
	}

	card := OrmModels.Card{}
	tx := db.Orm.First(&card, "id = ?", id)
	if tx.Error != nil {
		return false, gqlerror.Errorf("Error when get card in GetCard")
	}

	course := &OrmModels.Course{}
	tx = db.Orm.First(course, "id = ?", card.CourseID)
	if tx.Error != nil || tx.RowsAffected < 1 {
		return false, gqlerror.Errorf("Error when get course in GetCards")
	}
	if user.ID != course.UserID {
		return false, gqlerror.Errorf("User not allow to get card to this course")
	}

	return true, nil
}

// CreateCardsFromText is the resolver for the createCardsFromText field.
func (r *mutationResolver) CreateCardsFromText(ctx context.Context, input *model.NewCardInputFromText) ([]*model.Card, error) {
	user, err := GetUserFromContext(ctx)
	if err != nil {
		return nil, gqlerror.Errorf("Error when get user from context")
	}
	text := strings.Split(input.Text, "\n")
	var textResult [][]string

	for k, _ := range text {
		textResult = append(textResult, strings.Split(text[k], "\t"))
	}

	course := OrmModels.Course{
		Name:        input.Name,
		Description: *input.Description,
		UserID: user.ID,
	}
	var cards []OrmModels.Card

	tx := db.Orm.Create(&course)
	if tx.Error != nil {
		return nil, tx.Error
	}

	isError := false
	for _, v := range textResult {
		if len(v) == 2 {
			cards = append(cards, OrmModels.Card{
				Terminology: v[0],
				Definition:  v[1],
				CourseID:    course.ID,
			})
		} else {
			isError = true
			break
		}
	}
	if !isError {
		tx := db.Orm.Create(&cards)
		if tx.Error != nil {
			return nil, gqlerror.Errorf("Error when insert multiple cards to db, %v", tx.Error)
		}
	} else {
		return nil, gqlerror.Errorf("Input from clipboard is invalid")
	}

	var cardsGQL []*model.Card
	for _, v := range cards {
		cardsGQL = append(cardsGQL, &model.Card{
			ID:          v.ID,
			Terminology: &v.Terminology,
			Definition:  &v.Definition,
			CourseID:    v.CourseID,
		})
	}

	return cardsGQL, nil
}

// GetCards is the resolver for the getCards field.
func (r *queryResolver) GetCards(ctx context.Context, courseID *string) ([]*model.Card, error) {
	user, err := GetUserFromContext(ctx)
	if err != nil {
		return nil, gqlerror.Errorf("Error when get user from context")
	}

	course := &OrmModels.Course{}
	tx := db.Orm.First(course, "id = ?", courseID)
	if tx.Error != nil || tx.RowsAffected < 1 {
		return nil, gqlerror.Errorf("Error when get course in GetCards")
	}
	if user.ID != course.UserID {
		return nil, gqlerror.Errorf("User not allow to get cards to this course")
	}

	var cards []OrmModels.Card
	var cardsGQL []*model.Card
	tx = db.Orm.Where("course_id = ?", courseID).Find(&cards)
	fmt.Printf("%+v\n", cards)
	if tx.Error != nil {
		return nil, gqlerror.Errorf("Error when get cards in GetCards")
	}
	for k, _ := range cards {
		cardsGQL = append(cardsGQL, &model.Card{
			ID:          cards[k].ID,
			Terminology: &cards[k].Terminology,
			Definition:  &cards[k].Definition,
			CourseID:    cards[k].CourseID,
		})
	}

	return cardsGQL, nil
}

// GetCard is the resolver for the getCard field.
func (r *queryResolver) GetCard(ctx context.Context, id string) (*model.Card, error) {
	user, err := GetUserFromContext(ctx)
	if err != nil {
		return nil, gqlerror.Errorf("Error when get user from context")
	}

	card := OrmModels.Card{}
	tx := db.Orm.First(&card, "id = ?", id)
	if tx.Error != nil {
		return nil, gqlerror.Errorf("Error when get card in GetCard")
	}

	course := &OrmModels.Course{}
	tx = db.Orm.First(course, "id = ?", card.CourseID)
	if tx.Error != nil || tx.RowsAffected < 1 {
		return nil, gqlerror.Errorf("Error when get course in GetCards")
	}
	if user.ID != course.UserID {
		return nil, gqlerror.Errorf("User not allow to get card to this course")
	}
	cardORM := model.Card{
		ID:          card.ID,
		Terminology: &card.Terminology,
		Definition:  &card.Definition,
		CourseID:    card.CourseID,
	}

	return &cardORM, nil
}
