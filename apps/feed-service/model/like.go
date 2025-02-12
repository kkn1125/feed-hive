package model

type Like struct {
	Model
	UserId uint `gorm:"uniqueIndex:compositindex"`
	FeedId uint `gorm:"uniqueIndex:compositindex"`
}
