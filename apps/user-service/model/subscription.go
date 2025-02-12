package model

type Subscription struct {
	Model
	FollowerId  uint `gorm:"uniqueIndex:subscriberIndex"` // 구독한 유저
	FollowingId uint `gorm:"uniqueIndex:subscriberIndex"` // 구독할 대상 (피드 게시자)
	Follower    User `gorm:"foreignKey:FollowerId;constraint:onDelete:CASCADE,onUpdate:CASCADE;" json:"Follower"`
	Following   User `gorm:"foreignKey:FollowingId;constraint:onDelete:CASCADE,onUpdate:CASCADE;" json:"Following"`
}
