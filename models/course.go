package models

type Course struct {
	BaseUuid
	UserID string `json:"user_id"`
	Cards []Card
	Base
}
