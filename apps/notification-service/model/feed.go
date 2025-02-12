package model

type Feed struct {
	Model
	UserId  uint
	Content string

	Comments      []Comment      `gorm:"foreignKey:FeedId;joinForeignKey:FeedId;References:ID;joinReferences:FeedId" json:"Comments"`
	Notifications []Notification `gorm:"foreignKey:FeedId;joinForeignKey:FeedId;References:ID;joinReferences:FeedId" json:"Notifications"`
}
