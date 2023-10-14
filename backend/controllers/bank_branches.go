package controllers

import (
	_ "fmt"
	"github.com/bashkirian/gin-service/models"
	"github.com/bashkirian/gin-service/repo"
	"github.com/gin-gonic/gin"
	"net/http"
	_ "strconv"
	"database/sql"
)

// GET /branches
// Get all bank branches
func FindBanks(c *gin.Context) {
	// Get model if exist
	var banks []models.Bank
	var err error
	if banks, err = repo.GetBanks(c); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": banks})
}

// GET /branches/:id
// Find bank branch
func FindBank(c *gin.Context) {
	bank, err := repo.GetBank(c.Param("id"))
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": bank})
}

// // GET branches/optimal
// // Find nearest optimal branches
// // func FindBanks(c *gin.Context) {
// // 	// Get model if exist
// // 	var bank models.Bank
// // 	if err := models.DB.Where("id = ?", c.Param("id")).First(&bank).Error; err != nil {
// // 		c.JSON(http.StatusBadRequest, gin.H{"error": "Bank not found!"})
// // 		return
// // 	}

// // 	c.JSON(http.StatusOK, gin.H{"data": bank})
// // }

// //POST branches/:branchid/review
// //create review of bank branch
// func CreateReview(c *gin.Context) {
// 	// Validate input
// 	var input models.ReviewPost
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 	  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 	  return
// 	}

// 	// Create review
// 	bankid, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		panic("some error")
// 	}
// 	review := models.Review{Content: input.Content, BankID: bankid}
// 	models.DB.Create(&review)

// 	c.JSON(http.StatusOK, gin.H{"data": review})
//   }
