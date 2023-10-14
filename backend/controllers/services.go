package controllers

import (
	_ "fmt"
	"github.com/bashkirian/gin-service/models"
	"github.com/bashkirian/gin-service/repo"
	"github.com/gin-gonic/gin"
	"database/sql"
	"net/http"
	_ "strconv"
)

func InsertService(c *gin.Context) {
	var service models.Service
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	  	return
	}
	if err := repo.InsertService(c, service); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, c)
}