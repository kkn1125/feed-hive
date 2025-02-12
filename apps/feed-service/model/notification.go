package model

type Notification struct {
	Model
	FeedId     uint
	SenderId   uint // 피드 작성자
	ReceiverId uint // 알림 받을 사용자
	Message    string
	IsRead     bool `gorm:"default:false"`
}
