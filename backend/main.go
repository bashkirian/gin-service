package main

import (
	"fmt"
	"github.com/bashkirian/gin-service/controllers"
	"github.com/bashkirian/gin-service/models"
	"github.com/bashkirian/gin-service/repo"
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
		fmt.Println("cant parse port: %w", err)
	}
	err = models.PopulateDatabase()
	if err != nil {
		fmt.Println("populate: %w", err)
	}

	conf := repo.ConnectionConfig{
		Host:     os.Getenv("db_host"),
		Port:     port,
		User:     os.Getenv("db_user"),
		Password: os.Getenv("db_password"),
		DBName:   os.Getenv("db_name"),
	}
	if err = repo.ConnectDatabase(conf); err != nil {
		panic(err)
	}

	// Routes
	r.GET("/branches", controllers.FindBanks)
	r.GET("/branches/:id", controllers.FindBank)
	//r.POST("branches/:id/review", controllers.CreateReview)
	// Map
	r.GET("/map/route", controllers.FindRoute)
	r.GET("/clients/:id", controllers.GetClientInfo)
	r.POST("/clients/", controllers.InsertClient)
	// Run the server
	r.Run(":" + os.Getenv("app_port"))
}
