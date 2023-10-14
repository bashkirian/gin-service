package controllers

import (
	"fmt"
	"github.com/bashkirian/gin-service/models"
	"github.com/gin-gonic/gin"
	"net/http"
	_ "strconv"
	"github.com/gin-gonic/gin"
	"github.com/bashkirian/gin-service/models"
	"fmt"
	"database/sql"
)

// GET /branches
// Get all bank branches
func FindBanks(c *gin.Context) error {
	// Get model if exist
	var banks []*models.Bank
	rows, err := models.DB.Query("SELECT id, salepointname, latitude, longitude FROM bank.banks;")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return fmt.Errorf("%w", err)
	}
	for rows.Next() {
		b := new(models.Bank)
		err = rows.Scan(&b.ID, &b.Name, &b.Latitude, &b.Longitude)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return fmt.Errorf("rows scan: %w", err)
		}
		banks = append(banks, b)
	}
	if rows.Err() != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return fmt.Errorf("rows: %w", err)
	}
	c.JSON(http.StatusOK, gin.H{"data": banks})
	return nil
}

// GET /branches/:id
// Find bank branch
func FindBank(c *gin.Context) error {
	// Get model if exist
	var bank *models.Bank
	selectStatement := `SELECT id, salepointname, latitude, longitude FROM banks WHERE id = $1`
	err := models.DB.QueryRow(selectStatement, c.Param("id")).Scan(bank)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            return fmt.Errorf("bank %s: unknown album", c.Param("id"))
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": bank})
	return nil
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
