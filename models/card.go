package models

type Card struct {
	BaseUuid
	Terminology string `json:"termi,omitempty"`
	Definition  string `json:"defi,omitempty"`
	CourseID    string `json:"course_id"`
	Base
}
