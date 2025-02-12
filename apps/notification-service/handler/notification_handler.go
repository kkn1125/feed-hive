package handler

import (
	"feedhive/notifications/model"
	"feedhive/notifications/repository"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NotificationHandler struct {
	repo repository.NotificationRepository
}

func NewNotificationHandler(repo repository.NotificationRepository) *NotificationHandler {
	return &NotificationHandler{repo}
}

func (h *NotificationHandler) FindUnreadNotification(c *gin.Context) {
	userId := c.Param("userId")
	notifications, err := h.repo.FindUnread(userId)
	if err != nil {
		c.JSON(500, gin.H{
			"error": fmt.Sprintf("Failed to find notifications: %v", err),
		})
		return
	}

	c.JSON(200, notifications)
}

func (h *NotificationHandler) FindAllNotification(c *gin.Context) {
	notifications, err := h.repo.FindAll()
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to find notifications",
		})
		return
	}

	c.JSON(200, notifications)
}

func (h *NotificationHandler) FindNotificationById(c *gin.Context) {
	id := c.Param("id")
	notification, err := h.repo.FindById(id)
	if err != nil {
		c.JSON(404, gin.H{
			"error": "Failed to find notification",
		})
		return
	}

	c.JSON(200, notification)
}

func (handler *NotificationHandler) CreateNotification(c *gin.Context) {
	var notification model.Notification

	if err := c.ShouldBindJSON(&notification); err != nil {
		log.Println("Bind Error:", err)
		c.JSON(400, gin.H{
			"error": "Bind Error",
		})
		return
	}

	key, err := handler.repo.Create(&notification)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to create notification",
		})
		return
	}

	log.Println("key:", key)

	c.JSON(200, gin.H{
		"created": &notification.ID,
	})
}

func (handler *NotificationHandler) MarkAsRead(c *gin.Context) {
	notificationId := c.Param("notificationId")
	notificationIdUint, err := strconv.ParseUint(notificationId, 10, 32)

	if err != nil {
		log.Println("Invalid id", err)
		c.JSON(400, gin.H{
			"error": "Invalid id",
		})
		return
	}

	err = handler.repo.MarkAsRead(uint(notificationIdUint))

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to mark as read",
		})
		return
	}

	c.JSON(200, gin.H{
		"marked": &notificationId,
	})
}
