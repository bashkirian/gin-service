package controllers

import (
	"github.com/bashkirian/gin-service/models"
	"github.com/gin-gonic/gin"
	"net/http"
	_ "strconv"
)

// GET /branches
// Get all bank branches
func FindBanks(c *gin.Context) {
	// Get model if exist
	var banks []*models.Bank
	rows, err := models.DB.Query("SELECT id, salepointname, latitude, longitude FROM banks;")
	models.CheckError(err)
	for rows.Next() {
		b := new(models.Bank)
		err = rows.Scan(&b.ID, &b.Name, &b.Latitude, &b.Longitude)
		models.CheckError(err)
		banks = append(banks, b)
	}
	models.CheckError(rows.Err())
	c.JSON(http.StatusOK, gin.H{"data": banks})
}

// GET /branches/:id
// Find bank branch
func FindBank(c *gin.Context) {
	// Get model if exist
	var bank *models.Bank
	selectStatement := `SELECT id, salepointname, latitude, longitude FROM banks WHERE id = $1`
	rows, err := models.DB.Query(selectStatement, c.Param("id"))
	models.CheckError(err)
	for rows.Next() {
		b := new(models.Bank)
		err = rows.Scan(&b.ID, &b.Name, &b.Latitude, &b.Longitude)
		models.CheckError(err)
		bank = b
		models.CheckError(rows.Err())
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
