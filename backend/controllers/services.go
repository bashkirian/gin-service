package controllers

import (
	_ "fmt"
	"github.com/bashkirian/gin-service/models"
	"github.com/gin-gonic/gin"
	"net/http"
	_ "strconv"
)

func InsertService(c *gin.Context) {
	const query = `
INSERT INTO bank.bank_services VALUES ($1, $2);
`
	var service models.Service
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	  	return
	}
	if _, err := models.DB.ExecContext(c, query, service.ID, service.Description); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}