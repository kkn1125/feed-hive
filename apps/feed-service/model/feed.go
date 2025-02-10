package model

type Feed struct {
	Model
	UserId  uint
	Content string
	Comment []Comment `gorm:"constraint:onDelete:CASCADE,onUpdate:CASCADE;"`
}
