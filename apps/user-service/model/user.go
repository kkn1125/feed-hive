package model

type User struct {
	Model
	Name         string `gorm:"type:varchar(20)"`
	Email        string `gorm:"type:varchar(100)"`
	PasswordHash string `gorm:"type:varchar(200)" json:"-"`

	Followers  []Subscription `gorm:"foreignKey:FollowingId;joinForeignKey:FollowingId;References:ID;joinReferences:FollowingId" json:"Followers"`
	Followings []Subscription `gorm:"foreignKey:FollowerId;joinForeignKey:FollowerId;References:ID;joinReferences:FollowerId" json:"Followings"`
}

/* 수동 JSON 직렬화 */
// func (u *User) MarshalJSON() ([]byte, error) {
// 	type ReadUser struct {
// 		Model
// 		Name  string
// 		Email string
// 	}

// 	t := ReadUser{
// 		Model: u.Model,
// 		Name:  u.Name,
// 		Email: u.Email,
// 	}
// 	return json.Marshal(&t)
// }
