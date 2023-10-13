package main

import (
	"github.com/bashkirian/gin-service/controllers"
	"github.com/bashkirian/gin-service/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to database
	models.ConnectDatabase()
	models.MigrateDatabase()
	// Routes
	r.GET("/branches", controllers.FindBanks)
	r.GET("/branches/:id", controllers.FindBank)
	//r.POST("branches/:id/review", controllers.CreateReview)
	// Map
	r.GET("/map/route", controllers.FindRoute)
	// Run the server
	r.Run()
}
