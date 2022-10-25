package models

import (
	"fmt"
	"gorm.io/gorm"
)

func init() {
	fmt.Println("Init base models")
}

func AutoMigrate(DB *gorm.DB) error {
	return DB.AutoMigrate(User{})
}