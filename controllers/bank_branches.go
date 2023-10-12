package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/bashkirian/gin-service/models"
)

// GET /branches
// Get all bank branches 
func FindBanks(c *gin.Context) {
	var banks []models.Bank
	models.DB.Find(&banks)

	c.JSON(http.StatusOK, gin.H{"data": banks})
}

// GET /branches/:id
// Find bank branch
func FindBank(c *gin.Context) {
	// Get model if exist
	var bank models.Bank
	if err := models.DB.Where("id = ?", c.Param("branch_id")).First(&bank).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bank not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bank})
}

// GET branches/optimal
// Find nearest optimal branches
// func FindBanks(c *gin.Context) {
// 	// Get model if exist
// 	var bank models.Bank
// 	if err := models.DB.Where("id = ?", c.Param("id")).First(&bank).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Bank not found!"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": bank})
// }

//POST branches/:branchid/review
//create review of bank branch
func CreateReview(c *gin.Context) {
	// Validate input
	var input ReviewPost
	if err := c.ShouldBindJSON(&input); err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	  return
	}
  
	// Create review
	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)
  
	c.JSON(http.StatusOK, gin.H{"data": book})
  }