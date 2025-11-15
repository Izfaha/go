package initializers

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// host, user, password, dbname, port, sslmode string

func ConnectDB() {
	// initialize .env variables
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")

	openOrFail := func(dsn string) gorm.DB {
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			log.Fatal("[ERR] Failed to connect Database!")
		}
		return db
	}

}

func CreateDatabases(openOrFail func (dsn string), *gorm.DB) {
	adminDsn := fmt.Sprintf("postgres://%s:%s@%s:%s?sslmode=%s&search_path=public", user, password, host, port, sslmode)
	adminDB := openOrFail(adminDsn)
}
