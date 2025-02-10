package model

type Comment struct {
	Model
	UserId  uint
	FeedId  uint
	Content string
}
