package repository

import (
	"feedhive/users/database"
	"feedhive/users/model"
	"log"

	"gorm.io/gorm"
)

// func UserRepository() *gorm.DB {
// 	return database.DB
// }

type UserRepository interface {
	FindAll() (*[]model.User, error)
	FindById(id string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Create(user *model.User) (uint, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository() UserRepository {
	if database.DB == nil {
		log.Fatal("Database not connected")
	}
	return &userRepository{db: database.DB}
}

func (r *userRepository) FindAll() (*[]model.User, error) {
	var users []model.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

func (r *userRepository) FindById(id string) (*model.User, error) {
	var user model.User
	if err := r.db.Find(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	if err := r.db.Find(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Create(user *model.User) (uint, error) {
	result := r.db.Create(user)
	if result.Error != nil {
		return 0, result.Error
	}
	return user.ID, nil
}
