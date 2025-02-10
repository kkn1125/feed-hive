package main

import (
	"feedhive/feeds/database"
	"feedhive/feeds/model"
	"feedhive/feeds/router"
	"feedhive/feeds/variable"
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
		database.DB.Migrator().DropTable(&model.Feed{}, &model.Comment{}, &model.Likes{})
		database.DB.AutoMigrate(&model.Feed{}, &model.Comment{}, &model.Likes{})
	case "migrate":
		database.DB.AutoMigrate(&model.Feed{}, &model.Comment{}, &model.Likes{})
	default:
		r := setupRouter()
		r.Run(fmt.Sprintf(":%s", variable.PORT))
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())

	apiGuard := r.Group("/api")

	feedGroup := apiGuard.Group("/feeds")

	router.Feeds(feedGroup)

	return r
}
