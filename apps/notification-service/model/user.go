package model

type User struct {
	Model
	Name  string `gorm:"type:varchar(20)"`
	Email string `gorm:"type:varchar(100)"`
}
