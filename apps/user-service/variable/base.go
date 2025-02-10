package variable

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	err     error
	DB_USER string
	DB_PASS string
	DB_HOST string
	DB_PORT string
	DB_NAME string
	PORT    string
)

func LoadEnv() {
	mode := gin.Mode()
	if mode == gin.TestMode {
		// log.Fatal("Error")
		log.Println("Skip load env")
		return
	}

	err = godotenv.Load(".env")

	if err != nil {
		log.Println("Not found .env file")
	}

	DB_USER = os.Getenv("DB_USER")
	DB_PASS = os.Getenv("DB_PASS")
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_NAME = os.Getenv("DB_NAME")
	PORT = os.Getenv("PORT")
}
