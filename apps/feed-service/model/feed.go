package model

type Feed struct {
	Model
	UserId  uint
	Title   string
	Content string

	Comments []Comment `gorm:"foreignKey:FeedId;joinForeignKey:FeedId;References:ID;joinReferences:FeedId" json:"Comments"`

	Likes []Like `gorm:"foreignKey:FeedId;joinForeignKey:FeedId;References:ID;joinReferences:FeedId" json:"Likes"`
}
