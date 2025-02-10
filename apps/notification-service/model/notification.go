package model

type Notification struct {
	Model
	UserId  uint
	FeedId  uint
	Message string
	IsRead  bool `gorm:"default:false"`
}
