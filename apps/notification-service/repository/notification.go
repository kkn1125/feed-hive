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
	FindUnread(userId string) (*[]model.Notification, error)
	FindById(id string) (*model.Notification, error)
	Create(notification *model.Notification) (uint, error)
	CreateFeedNotification(feedId uint) (uint, error)
	MarkAsRead(notificationId uint) error
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

func (r *notificationRepository) FindUnread(userId string) (*[]model.Notification, error) {
	var notification []model.Notification
	if err := r.db.Where("is_read = ? AND receiver_id = ?", false, userId).Find(&notification, userId).Error; err != nil {
		return nil, err
	}
	return &notification, nil
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
	notification.SenderId = feed.UserId

	users, err := r.GetAllUsersExcept(feed.UserId)

	if err != nil {
		log.Printf("CreateFeedNotification Error: %v\n", err)
		return 0, err
	}

	for _, user := range *users {
		notification.ReceiverId = user.ID
		notification.Message = "new feed"
		if err := r.db.Create(&notification).Error; err != nil {
			log.Printf("CreateFeedNotification Error: %v\n", err)
			return 0, err
		}
	}

	// result := r.db.Create(&notification)
	// if result.Error != nil {
	// 	return 0, result.Error
	// }
	return notification.ID, nil
}

func (r *notificationRepository) GetAllUsersExcept(senderId uint) (*[]model.User, error) {
	var users []model.User
	if err := r.db.Where("id != ?", senderId).Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

func (r *notificationRepository) MarkAsRead(notificationId uint) error {
	notification := model.Notification{}
	if err := r.db.Model(&notification).Where("id = ?", notificationId).Update("is_read", true).Error; err != nil {
		return err
	}
	return nil
}
