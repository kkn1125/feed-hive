package repository

import (
	"feedhive/users/database"
	"feedhive/users/model"
	"fmt"
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
	Subscribe(followerId uint, followingId uint) (bool, error)
	GetSubscriptions(followerId uint) (*[]model.User, error)
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

func (r *userRepository) GetSubscriptions(followerId uint) (*[]model.User, error) {
	var subscriptions []model.User

	if err := r.db.Raw(`
		SELECT
			u.id, u.name, u.email,
			u.created_at, u.updated_at, u.deleted_at
		FROM users u
		WHERE u.id = ?
	`, followerId).Scan(&subscriptions).Error; err != nil {
		log.Printf("GetSubscriptions Error: %v\n", err)
		return nil, err
	}

	for i := range subscriptions {
		if err := r.db.Model(&subscriptions[i]).Association("Followers").Find(&subscriptions[i].Followers); err != nil {
			log.Printf("Failed to load Followings: %v\n", err)
		}
		if err := r.db.Model(&subscriptions[i]).Association("Followings").Find(&subscriptions[i].Followings); err != nil {
			log.Printf("Failed to load Followers: %v\n", err)
		}
	}

	return &subscriptions, nil
}

func (r *userRepository) Subscribe(followerId uint, followingId uint) (bool, error) {
	result := r.db.Find(&model.User{}, followerId)

	if result.Error != nil {
		log.Printf("Subscribe Error: %v\n", result.Error)
		return false, result.Error
	}

	if result.RowsAffected == 0 {
		message := fmt.Errorf("not found follower_id: %v", followerId)
		return false, message
	}

	result = r.db.Find(&model.User{}, followingId)

	if result.Error != nil {
		log.Printf("Subscribe Error: %v\n", result.Error)
		return false, result.Error
	}

	if result.RowsAffected == 0 {
		message := fmt.Errorf("not found following_id: %v", followingId)
		return false, message
	}

	result = r.db.Create(&model.Subscription{FollowerId: followerId, FollowingId: followingId})

	if result.Error != nil {
		log.Printf("Subscribe Error: %v\n", result.Error)
		return false, result.Error
	}
	return true, nil
}
