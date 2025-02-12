package handler

import (
	"feedhive/users/model"
	"feedhive/users/repository"
	"feedhive/users/util"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	repo repository.UserRepository
}

func NewUserHandler(repo repository.UserRepository) *UserHandler {
	return &UserHandler{repo}
}

func (h *UserHandler) FindAllUser(c *gin.Context) {
	users, err := h.repo.FindAll()
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to find users",
		})
		return
	}

	c.JSON(200, users)
}

func (h *UserHandler) FindUserById(c *gin.Context) {
	id := c.Param("id")
	user, err := h.repo.FindById(id)
	if err != nil {
		c.JSON(404, gin.H{
			"error": "Failed to find user",
		})
		return
	}

	c.JSON(200, user)
}

func (h *UserHandler) FindUserByEmail(c *gin.Context) {
	email := c.Param("email")
	user, err := h.repo.FindByEmail(email)
	if err != nil {
		c.JSON(404, gin.H{
			"error": "Failed to find user",
		})
		return
	}
	c.JSON(200, user)
}

func (handler *UserHandler) CreateUser(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println("Bind Error:", err)
		c.JSON(400, gin.H{
			"error": "Bind Error",
		})
		return
	}

	user.PasswordHash = util.HashPassword(user.PasswordHash)

	key, err := handler.repo.Create(&user)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	log.Println("key:", key)

	c.JSON(200, gin.H{
		"created": &user.ID,
	})
}

func (handler *UserHandler) Subscribe(c *gin.Context) {
	followerId := c.Param("followerId")
	followingId := c.Param("followingId")

	if followerId == followingId {
		c.JSON(400, gin.H{
			"error": "Follower and following id cannot be the same",
		})
		return
	}

	followerIdUint, err := strconv.ParseUint(followerId, 10, 32)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid id",
		})
		return
	}

	followingIdUint, err := strconv.ParseUint(followingId, 10, 32)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid id",
		})
		return
	}

	isSubscribed, err := handler.repo.Subscribe(uint(followerIdUint), uint(followingIdUint))
	if err != nil {
		c.JSON(500, gin.H{
			"subscribed": isSubscribed,
			"message":    err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"subscribed": isSubscribed,
	})
}

func (handler *UserHandler) GetSubscriptions(c *gin.Context) {
	followerId := c.Param("followerId")

	followerIdUint, err := strconv.ParseUint(followerId, 10, 32)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid id",
		})
		return
	}

	subscriptions, err := handler.repo.GetSubscriptions(uint(followerIdUint))
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, subscriptions)
}
