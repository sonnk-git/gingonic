package models

type Course struct {
	BaseUuid
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      string `json:"user_id"`
	Cards       []Card
	Base
}
