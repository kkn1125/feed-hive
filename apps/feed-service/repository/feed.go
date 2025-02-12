package repository

import (
	"feedhive/feeds/database"
	"feedhive/feeds/model"
	"log"

	"gorm.io/gorm"
)

// func FeedRepository() *gorm.DB {
// 	return database.DB
// }

type FeedRepository interface {
	FindAll() (*[]model.Feed, error)
	FindById(id string) (*model.Feed, error)
	FindNotificationById(notificationId string) (*model.Notification, error)
	Create(feed *model.Feed) (uint, error)
}

type feedRepository struct {
	db *gorm.DB
}

func NewFeedRepository() FeedRepository {
	if database.DB == nil {
		log.Fatal("Database not connected")
	}
	return &feedRepository{db: database.DB}
}

func (r *feedRepository) FindAll() (*[]model.Feed, error) {
	var feeds []model.Feed
	if err := r.db.Find(&feeds).Error; err != nil {
		return nil, err
	}
	return &feeds, nil
}

func (r *feedRepository) FindById(id string) (*model.Feed, error) {
	var feed model.Feed
	if err := r.db.Find(&feed, id).Error; err != nil {
		return nil, err
	}
	return &feed, nil
}

func (r *feedRepository) FindNotificationById(notificationId string) (*model.Notification, error) {
	var notification model.Notification
	if err := r.db.Find(&notification, notificationId).Error; err != nil {
		return nil, err
	}
	return &notification, nil
}

func (r *feedRepository) Create(feed *model.Feed) (uint, error) {
	result := r.db.Create(feed)
	if result.Error != nil {
		return 0, result.Error
	}
	return feed.ID, nil
}
