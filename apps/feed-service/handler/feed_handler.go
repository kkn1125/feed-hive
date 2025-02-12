package handler

import (
	"feedhive/feeds/library"
	"feedhive/feeds/model"
	"feedhive/feeds/repository"
	"log"

	"github.com/gin-gonic/gin"
)

type FeedHandler struct {
	repo repository.FeedRepository
}

func NewFeedHandler(repo repository.FeedRepository) *FeedHandler {
	return &FeedHandler{repo}
}

func (h *FeedHandler) FindAllFeed(c *gin.Context) {
	feeds, err := h.repo.FindAll()
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to find feeds",
		})
		return
	}

	c.JSON(200, feeds)
}

func (h *FeedHandler) FindFeedById(c *gin.Context) {
	id := c.Param("id")
	feed, err := h.repo.FindById(id)
	if err != nil {
		c.JSON(404, gin.H{
			"error": "Failed to find feed",
		})
		return
	}

	c.JSON(200, feed)
}

func (handler *FeedHandler) CreateFeed(c *gin.Context) {
	var feed model.Feed

	if err := c.ShouldBindJSON(&feed); err != nil {
		log.Println("Bind Error:", err)
		c.JSON(400, gin.H{
			"error": "Bind Error",
		})
		return
	}
	log.Printf("feed: %v, content: %v", feed.UserId, feed.Content)
	key, err := handler.repo.Create(&feed)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to create feed",
		})
		return
	}

	log.Println("key:", key)

	go library.Send(&feed)

	c.JSON(200, gin.H{
		"created": &feed.ID,
	})
}

// func (handler *FeedHandler) SendMarkAsRead(c *gin.Context) {
// 	notificationId := c.Param("notificationId")
// 	notification, err := handler.repo.FindNotificationById(notificationId)
// 	if err != nil {
// 		c.JSON(404, gin.H{
// 			"error": "Failed to find feed",
// 		})
// 		return
// 	}

// 	go library.SendMarkAsRead(notification.ID)

// 	log.Printf("notification: %v, content: %v", notification.ID, notification.Message)

// 	c.JSON(200, gin.H{
// 		"readed": &notification.ID,
// 	})

// }
