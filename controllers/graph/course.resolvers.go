package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	model "gingonic/graph"
)

// CreateCourse is the resolver for the createCourse field.
func (r *mutationResolver) CreateCourse(ctx context.Context, input model.NewCourseInput) (*model.Course, error) {
	panic(fmt.Errorf("not implemented: CreateCourse - createCourse"))
}

// EditCourse is the resolver for the editCourse field.
func (r *mutationResolver) EditCourse(ctx context.Context, input model.CourseInput) (*model.Course, error) {
	panic(fmt.Errorf("not implemented: EditCourse - editCourse"))
}

// DeleteCourse is the resolver for the deleteCourse field.
func (r *mutationResolver) DeleteCourse(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteCourse - deleteCourse"))
}

// GetCourses is the resolver for the getCourses field.
func (r *queryResolver) GetCourses(ctx context.Context) ([]*model.Course, error) {
	panic(fmt.Errorf("not implemented: GetCourses - getCourses"))
}

// GetCourse is the resolver for the getCourse field.
func (r *queryResolver) GetCourse(ctx context.Context, id string) (*model.Course, error) {
	panic(fmt.Errorf("not implemented: GetCourse - getCourse"))
}
