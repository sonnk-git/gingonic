package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"gingonic/db"
	model "gingonic/graph"
	OrmModels "gingonic/models"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

// SetSubscribe is the resolver for the setSubscribe field.
func (r *mutationResolver) SetSubscribe(ctx context.Context, input model.NotificationRequest) (*model.Notification, error) {
	user, err := GetUserFromContext(ctx)
	if err != nil {
		return nil, gqlerror.Errorf("Error when get user from context")
	}

	var sub OrmModels.Subscription
	tx := db.Orm.Find(&sub, "user_id = ?", user.ID)
	if tx.Error != nil {
		return nil, gqlerror.Errorf("Error when find subscription in db")
	}

	sub.EveryMinute = 1

	if input.SubscribeState != nil {
		sub.SubscribeState = *input.SubscribeState
	} else {
		sub.SubscribeState = false
	}

	if input.Sub != nil {
		sub.Sub = *input.Sub
	}

	if input.CourseID != nil {
		sub.CourseID = *input.CourseID
	}

	if tx.RowsAffected > 0 {
		tx := db.Orm.Save(&sub)
		if tx.Error != nil {
			return nil, gqlerror.Errorf("Error when save subscription in db")
		}
	} else {
		sub.UserID = user.ID
		tx := db.Orm.Create(&sub)
		if tx.Error != nil {
			return nil, gqlerror.Errorf("Error when create new subscription in db")
		}
	}

	return convertNotificationORMToGQL(sub), nil
}

// GetSubscription is the resolver for the getSubscription field.
func (r *queryResolver) GetSubscription(ctx context.Context) (*model.Notification, error) {
	user, err := GetUserFromContext(ctx)
	if err != nil {
		return nil, gqlerror.Errorf("Error when get user from context")
	}

	var sub OrmModels.Subscription
	tx := db.Orm.Find(&sub, "user_id = ?", user.ID)
	if tx.Error != nil {
		return nil, gqlerror.Errorf("Error when find subscription in db")
	}
	return convertNotificationORMToGQL(sub), nil
}

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func convertNotificationORMToGQL(subscription OrmModels.Subscription) *model.Notification {
	return &model.Notification{
		ID:             subscription.ID,
		UserID:         subscription.UserID,
		CourseID:       &subscription.CourseID,
		Sub:            &subscription.Sub,
		SubscribeState: &subscription.SubscribeState,
		EveryMinute:    &subscription.EveryMinute,
	}
}
