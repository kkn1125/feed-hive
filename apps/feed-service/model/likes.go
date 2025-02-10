package model

type Likes struct {
	Model
	UserId uint `gorm:"uniqueIndex:compositindex"`
	FeedId uint `gorm:"uniqueIndex:compositindex"`
}
