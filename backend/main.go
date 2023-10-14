package main

import (
	"fmt"
	"github.com/bashkirian/gin-service/controllers"
	"github.com/bashkirian/gin-service/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	fmt.Println("OK")

	// Connect to database
	models.ConnectDatabase(models.ConnectionConfig{})
	models.MigrateDatabase()
	// Routes
	// Banks
	r.GET("/branches", controllers.FindBanks)
	r.GET("/branches/:id", controllers.FindBank)
	//r.POST("branches/:id/review", controllers.CreateReview)
	// Services
	r.POST("/services", controllers.InsertService)
	// Map
	r.GET("/map/route", controllers.FindRoute)
	// Run the server
	r.Run()
}
