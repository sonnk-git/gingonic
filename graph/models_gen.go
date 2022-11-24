// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Card struct {
	ID          string  `json:"id"`
	Terminology *string `json:"terminology"`
	Definition  *string `json:"definition"`
	CourseID    string  `json:"courseId"`
}

type CardInput struct {
	ID          string  `json:"id"`
	Terminology *string `json:"terminology"`
	Definition  *string `json:"definition"`
}

type Course struct {
	ID          string `json:"id"`
	UserID      string `json:"userId"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CourseInput struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type NewCardInput struct {
	CourseID    string  `json:"courseId"`
	Terminology *string `json:"terminology"`
	Definition  *string `json:"definition"`
}

type NewCourseInput struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type NewUser struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
