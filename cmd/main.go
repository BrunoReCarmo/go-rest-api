package main

import (
	"github.com/gin-gonic/gin"
	"go-api/controller"
)

func main() {
	// Server is using the Gin package
	server := gin.Default()

	//Controllers
	OfferController := controller.NewOfferController()

	// The endpoint will be "/hello"
	// Simple inline function, gin.H maps the response message
	// The expected response is:

	server.GET("/offers", OfferController.GetOffers)
	
	server.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "world",
		})
	})

	//localhost:4000
	server.Run(":4000")
}
