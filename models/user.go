package models

import (
	"fmt"
	"gorm.io/gorm"
)

func init() {
	fmt.Println("Init user models")
}

type User struct {
	gorm.Model
	Name  string
	Email string
	Password string
}
