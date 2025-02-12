package model

type Comment struct {
	Model
	UserId  uint
	FeedId  uint
	Content string

	Feed Feed `gorm:"foreignKey:FeedId;constraint:onDelete:CASCADE,onUpdate:CASCADE;"`
}
