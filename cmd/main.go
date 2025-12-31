package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Server is using the Gin package
	server := gin.Default()

	// The endpoint will be "/hello"
	// Simple inline function, gin.H maps the response message
	// The expected response is:
	
	server.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "world",
		})
	})

	//localhost:4000
	server.Run(":4000")
}
