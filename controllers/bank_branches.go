package controllers

import (
	"net/http"
	"strconv"
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
	var input models.ReviewPost
	if err := c.ShouldBindJSON(&input); err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	  return
	}
  
	// Create review
	bankid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic("some error")
	}
	review := models.Review{Content: input.Content, BankID: bankid}
	models.DB.Create(&review)
  
	c.JSON(http.StatusOK, gin.H{"data": review})
  }