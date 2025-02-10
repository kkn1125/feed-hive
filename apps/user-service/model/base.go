package model

import (
	"feedhive/users/util"
)

type Model struct { // size=88 (0x58)
	ID        uint `gorm:"primarykey"`
	CreatedAt util.DateTime
	UpdatedAt util.DateTime
	DeletedAt util.NullTime `gorm:"index"`
}
