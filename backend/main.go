package main

import (
	"fmt"
	"github.com/bashkirian/gin-service/controllers"
	"github.com/bashkirian/gin-service/models"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
)

func main() {
	r := gin.Default()

	fmt.Println("OK")

	// Connect to database
	port, err := strconv.Atoi(os.Getenv("db_port"))
	if err != nil {
		print(fmt.Errorf("cant parse port: %w", err))
	}
	err = models.PopulateDatabase()
	if err != nil {
		print(fmt.Errorf("populate: %w", err))
	}
	if err == nil {
		print("OKOKOKOK")
	}
	err = models.ConnectDatabase(models.ConnectionConfig{
		Host:     os.Getenv("db_host"),
		Port:     port,
		User:     os.Getenv("db_user"),
		Password: os.Getenv("db_password"),
		DBName:   os.Getenv("db_name"),
	})
	if err != nil {
		print(err)
	}
	// Routes
	r.GET("/branches", controllers.FindBanks)
	r.GET("/branches/:id", controllers.FindBank)
	//r.POST("branches/:id/review", controllers.CreateReview)
	// Map
	r.GET("/map/route", controllers.FindRoute)
	// Run the server
	r.Run(":" + os.Getenv("app_port"))
}
