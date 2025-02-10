package handler

import (
	"feedhive/notifications/model"
	"feedhive/notifications/repository"
	"log"

	"github.com/gin-gonic/gin"
)

type NotificationHandler struct {
	repo repository.NotificationRepository
}

func NewNotificationHandler(repo repository.NotificationRepository) *NotificationHandler {
	return &NotificationHandler{repo}
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
