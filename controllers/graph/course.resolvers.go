package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"gingonic/db"
	model "gingonic/graph"
	OrmModels "gingonic/models"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

// CreateCourse is the resolver for the createCourse field.
func (r *mutationResolver) CreateCourse(ctx context.Context, input model.NewCourseInput) (*model.Course, error) {
	user, err := GetUserFromContext(ctx)
	if err != nil {
		return nil, gqlerror.Errorf("Error when get user from context")
	}

	course := &OrmModels.Course{
		UserID:      user.ID,
		Name:        input.Name,
		Description: *input.Description,
	}

	tx := db.Orm.Create(&course)
	if tx.Error != nil {
		return nil, gqlerror.Errorf("Error when create course in db")
	}

	courseGQL := &model.Course{
		ID:          course.ID,
		UserID:      course.UserID,
		Name:        course.Name,
		Description: course.Description,
	}

	return courseGQL, nil
}

// EditCourse is the resolver for the editCourse field.
func (r *mutationResolver) EditCourse(ctx context.Context, input model.CourseInput) (*model.Course, error) {
	user, err := GetUserFromContext(ctx)
	if err != nil {
		return nil, gqlerror.Errorf("Error when get user from context")
	}

	course := OrmModels.Course{}
	tx := db.Orm.First(&course, "id = ?", input.ID)
	if tx.Error != nil {
		return nil, gqlerror.Errorf("Error when get course from context")
	}

	if course.UserID != user.ID {
		return nil, gqlerror.Errorf("User is not owner of course")
	}

	tx = db.Orm.Model(&course).Updates(OrmModels.Course{
		Name:        input.Name,
		Description: input.Description,
	})
	if tx.Error != nil {
		return nil, gqlerror.Errorf("Error when update course")
	}

	courseGQL := &model.Course{
		ID:          course.ID,
		UserID:      course.UserID,
		Name:        course.Name,
		Description: course.Description,
	}

	return courseGQL, nil
}

// DeleteCourse is the resolver for the deleteCourse field.
func (r *mutationResolver) DeleteCourse(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteCourse - deleteCourse"))
}

// GetCourses is the resolver for the getCourses field.
func (r *queryResolver) GetCourses(ctx context.Context) ([]*model.Course, error) {
	user, err := GetUserFromContext(ctx)
	if err != nil {
		return nil, gqlerror.Errorf("Error when get user from context")
	}

	var courses []OrmModels.Course
	db.Orm.Where("user_id = ?", user.ID).Find(&courses)
	coursesGQL := make([]*model.Course, 0, len(courses))

	for _, v := range courses {
		coursesGQL = append(coursesGQL, &model.Course{
			ID:          v.ID,
			UserID:      v.UserID,
			Name:        v.Name,
			Description: v.Description,
		})
	}

	return coursesGQL, nil
}

// GetCourse is the resolver for the getCourse field.
func (r *queryResolver) GetCourse(ctx context.Context, id string) (*model.Course, error) {
	user, err := GetUserFromContext(ctx)
	if err != nil {
		return nil, gqlerror.Errorf("Error when get user from context")
	}

	course := &OrmModels.Course{}
	tx := db.Orm.First(&course, "user_id = ? AND id = ?", user.ID, id)
	if tx.Error != nil || tx.RowsAffected < 1 {
		return nil, gqlerror.Errorf("Error when get course from GetCourse")
	}
	courseGQL := model.Course{
		ID:          course.ID,
		UserID:      course.UserID,
		Name:        course.Name,
		Description: course.Description,
	}

	return &courseGQL, nil
}
