package main

import (
	"jwt-go/controllers"
	"jwt-go/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	initializers.SyncDatabases()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Singup)
	r.Run()
}
