package database

import (
	"feedhive/users/variable"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var newLogger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	logger.Config{
		SlowThreshold:             time.Second,   // Slow SQL threshold
		LogLevel:                  logger.Silent, // Log level
		IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
		ParameterizedQueries:      true,          // Don't include params in the SQL log
		Colorful:                  false,         // Disable color
	},
)

var DB *gorm.DB

func Connect() {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		variable.DB_USER,
		variable.DB_PASS,
		variable.DB_HOST,
		variable.DB_PORT,
		variable.DB_NAME,
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
	})
	if err != nil {
		log.Fatal("failed to connect database")
	}
}
