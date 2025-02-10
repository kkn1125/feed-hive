package router

import (
	"feedhive/feeds/handler"
	"feedhive/feeds/repository"

	"github.com/gin-gonic/gin"
)

// var feedRepo repository.FeedRepository
// var feedHandler *handler.FeedHandler

func Feeds(r *gin.RouterGroup) {
	feedRepo := repository.NewFeedRepository()
	feedHandler := handler.NewFeedHandler(feedRepo)

	r.GET("/", feedHandler.FindAllFeed)
	r.GET("/:id", feedHandler.FindFeedById)
	r.POST("/", feedHandler.CreateFeed)
}
