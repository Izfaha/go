package initializers

import (
	"os"

	"gorm.io/gorm"
)

var DB *gorm.DB
	// host, user, password, dbname, port, sslmode string

func ConnectDB() {
	// initialize .env variables
	// host := os.Getenv("DB_HOST")
	// user := os.Getenv("DB_USER")
	// password := os.Getenv("DB_PASSWORD")
	// dbname := os.Getenv("DB_NAME")
	// port := os.Getenv("DB_PORT")
	// sslmode := os.Getenv("DB_SSLMODE")
	var err error

	dsn := os.Getenv("DB")
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("[ERR] Failed to connect Database!")
	}
}
