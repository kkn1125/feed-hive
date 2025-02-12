package router

import (
	"feedhive/notifications/handler"
	"feedhive/notifications/repository"

	"github.com/gin-gonic/gin"
)

func Notifications(r *gin.RouterGroup) {
	notificationRepo := repository.NewNotificationRepository()
	notificationHandler := handler.NewNotificationHandler(notificationRepo)

	r.GET("/", notificationHandler.FindAllNotification)
	r.GET("/user/:userId/unread", notificationHandler.FindUnreadNotification)
	r.GET("/:id", notificationHandler.FindNotificationById)
	r.POST("/:notificationId/read", notificationHandler.MarkAsRead)
	r.POST("/", notificationHandler.CreateNotification)
}
