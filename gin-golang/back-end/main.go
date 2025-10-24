package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	// endpoint
	router.GET("/data", func(ctx *gin.Context){
		ctx.JSON(200, gin.H{
			"message": "OK!!",
		})
	})
	router.Run(":5869")
}
