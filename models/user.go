package models

import (
	"fmt"
)

func init() {
	fmt.Println("Init user models")
}

type User struct {
	BaseUuid
	Name           string `json:"name,omitempty"`
	Email          string `gorm:"unique" json:"email,omitempty" validate:"required,email"`
	Password       string `json:"password,omitempty" validate:"required"`
	Subscription   string `json:"subscription,omitempty"`
	SubscribeState bool   `json:"subscribe_state,omitempty"`
	Courses        []Course
	Base
}
