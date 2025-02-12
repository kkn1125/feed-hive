package router

import (
	"feedhive/users/handler"
	"feedhive/users/repository"

	"github.com/gin-gonic/gin"
)

// var userRepo repository.UserRepository
// var userHandler *handler.UserHandler

func Users(r *gin.RouterGroup) {
	userRepo := repository.NewUserRepository()
	userHandler := handler.NewUserHandler(userRepo)

	r.GET("/", userHandler.FindAllUser)
	r.GET("/:id", userHandler.FindUserById)
	r.GET("/email/:email", userHandler.FindUserByEmail)
	r.GET("/subscriptions/:followerId", userHandler.GetSubscriptions)
	r.POST("/", userHandler.CreateUser)
	r.POST("/follow/:followerId/:followingId", userHandler.Subscribe)
}
