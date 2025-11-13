package initializers

import "jwt-go/models"

func SyncDatabases() {
	// migrate tables
	DB.AutoMigrate(&models.User{})
}