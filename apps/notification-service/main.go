package main

import (
	"feedhive/notifications/database"
	"feedhive/notifications/library"
	"feedhive/notifications/model"
	"feedhive/notifications/router"
	"feedhive/notifications/variable"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	variable.LoadEnv()
	database.Connect()

	arg := ""

	args := os.Args
	if len(args) > 1 {
		arg = args[1]
		log.Println(arg)
	}

	switch arg {
	case "migrateforce":
		database.DB.Migrator().DropTable(&model.Notification{})
		database.DB.AutoMigrate(&model.Notification{})
	case "migrate":
		database.DB.AutoMigrate(&model.Notification{})
	default:
		go library.Receive()
		// go library.ReceiveMarkAsRead()
		r := setupRouter()
		r.Run(fmt.Sprintf(":%s", variable.PORT))
	}

}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())

	apiGuard := r.Group("/api")

	notificationGroup := apiGuard.Group("/notifications")

	router.Notifications(notificationGroup)

	return r
}
