package main

import (
	"feedhive/users/database"
	"feedhive/users/model"
	"feedhive/users/router"
	"feedhive/users/variable"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

// var db = make(map[string]string)

// func setupRouter() *gin.Engine {
// 	// Disable Console Color
// 	// gin.DisableConsoleColor()
// 	r := gin.Default()

// 	// Ping test
// 	r.GET("/ping", func(c *gin.Context) {
// 		c.String(http.StatusOK, "pong")
// 	})

// 	// Get user value
// 	r.GET("/user/:name", func(c *gin.Context) {
// 		user := c.Params.ByName("name")
// 		value, ok := db[user]
// 		if ok {
// 			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
// 		} else {
// 			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
// 		}
// 	})

// 	// Authorized group (uses gin.BasicAuth() middleware)
// 	// Same than:
// 	// authorized := r.Group("/")
// 	// authorized.Use(gin.BasicAuth(gin.Credentials{
// 	//	  "foo":  "bar",
// 	//	  "manu": "123",
// 	//}))
// 	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
// 		"foo":  "bar", // user:foo password:bar
// 		"manu": "123", // user:manu password:123
// 	}))

// 	/* example curl for /admin with basicauth header
// 	   Zm9vOmJhcg== is base64("foo:bar")

// 		curl -X POST \
// 	  	http://localhost:8080/admin \
// 	  	-H 'authorization: Basic Zm9vOmJhcg==' \
// 	  	-H 'content-type: application/json' \
// 	  	-d '{"value":"bar"}'
// 	*/
// 	authorized.POST("admin", func(c *gin.Context) {
// 		user := c.MustGet(gin.AuthUserKey).(string)

// 		// Parse JSON
// 		var json struct {
// 			Value string `json:"value" binding:"required"`
// 		}

// 		if c.Bind(&json) == nil {
// 			db[user] = json.Value
// 			c.JSON(http.StatusOK, gin.H{"status": "ok"})
// 		}
// 	})

// 	return r
// }

// func main() {
// 	r := setupRouter()
// 	// Listen and Server in 0.0.0.0:8080
// 	r.Run(":8080")
// }

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
		database.DB.Migrator().DropTable(&model.User{}, &model.Subscription{})
		database.DB.AutoMigrate(&model.User{}, &model.Subscription{})
	case "migrate":
		database.DB.AutoMigrate(&model.User{}, &model.Subscription{})
	default:
		r := setupRouter()
		r.Run(fmt.Sprintf(":%s", variable.PORT))
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())

	apiGuard := r.Group("/api")

	userGroup := apiGuard.Group("/users")

	router.Users(userGroup)

	return r
}
