package models

import (
	"fmt"
)

func init() {
	fmt.Println("Init user models")
}

type User struct {
	BaseUuid
	Name     string `json:"name,omitempty"`
	Email    string `gorm:"unique" json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Courses  []Course
	Base
}
