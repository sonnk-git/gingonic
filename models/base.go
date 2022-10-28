package models

import (
	"fmt"
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
	"time"
)

func init() {
	fmt.Println("Init base models")
}

type BaseUuid struct {
	ID string `gorm:"primaryKey" json:"id"`
}

type Base struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (b *BaseUuid) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = ulid.Make().String()
	return
}

func AutoMigrate(DB *gorm.DB) error {
	return DB.AutoMigrate(User{})
}
