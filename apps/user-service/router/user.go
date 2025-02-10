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
	r.POST("/", userHandler.CreateUser)
}
