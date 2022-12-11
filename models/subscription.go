package models

type Subscription struct {
	BaseUuid
	UserID         string `json:"user_id"`
	CourseID       string `json:"course_id"`
	Sub            string `json:"sub,omitempty"`
	SubscribeState bool   `gorm:"default:false" json:"subscribe_state,omitempty"`
	EveryMinute    int    `json:"every_minute,omitempty"`
	Base
}
