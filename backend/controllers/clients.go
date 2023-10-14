package controllers

import (
	"github.com/bashkirian/gin-service/models"
	"github.com/bashkirian/gin-service/repo"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func GetClientInfo(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "incorrect id")
		return
	}

	la, lo, err := repo.GetClientLocationByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, "not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": []string{la, lo}})
}

func InsertClient(c *gin.Context) {
	type request struct {
		Lo string `json:"lo"`
		La string `json:"la"`
	}
	var req request
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"": err.Error()})
		return
	}

	err = repo.InsertClient(c, models.Client{
		ID:          uuid.New(),
		LocationLon: req.Lo,
		LocationLat: req.La,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"": err.Error()})
		return
	}

}
