package repository

import (
	"feedhive/notifications/database"
	"feedhive/notifications/model"
	"log"

	"gorm.io/gorm"
)

// func NotificationRepository() *gorm.DB {
// 	return database.DB
// }

type NotificationRepository interface {
	FindAll() (*[]model.Notification, error)
	FindById(id string) (*model.Notification, error)
	Create(notification *model.Notification) (uint, error)
	CreateFeedNotification(feedId uint) (uint, error)
}

type notificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository() NotificationRepository {
	if database.DB == nil {
		log.Fatal("Database not connected")
	}
	return &notificationRepository{db: database.DB}
}

func (r *notificationRepository) FindAll() (*[]model.Notification, error) {
	var notifications []model.Notification
	if err := r.db.Find(&notifications).Error; err != nil {
		return nil, err
	}
	return &notifications, nil
}

func (r *notificationRepository) FindById(id string) (*model.Notification, error) {
	var notification model.Notification
	if err := r.db.Find(&notification, id).Error; err != nil {
		return nil, err
	}
	return &notification, nil
}

func (r *notificationRepository) Create(notification *model.Notification) (uint, error) {
	result := r.db.Create(notification)
	if result.Error != nil {
		return 0, result.Error
	}
	return notification.ID, nil
}

func (r *notificationRepository) CreateFeedNotification(feedId uint) (uint, error) {
	var feed model.Feed
	if err := r.db.First(&feed, feedId).Error; err != nil {
		return 0, err
	}
	var notification model.Notification
	notification.FeedId = feedId
	notification.UserId = feed.UserId
	result := r.db.Create(&notification)
	if result.Error != nil {
		return 0, result.Error
	}
	return notification.ID, nil
}
