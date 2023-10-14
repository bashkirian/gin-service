package controllers

import (
	_ "fmt"
	"github.com/bashkirian/gin-service/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
	"net/http"
	_ "strconv"
)

const apiKey = "AIzaSyDZ-FFbuQ0xhk1ArrZW8zZ8LdpUuIDsD0g"

// GET /map/route
// Get route from two points and distance
func FindRoute(c *gin.Context) {
	var points models.MapPointPayload
	if err := c.ShouldBindJSON(&points); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"map json binding": err.Error()})
		return
	}
	origin := points.MapPoints[0].Latitude
	origin += ", "
	origin += points.MapPoints[0].Longtitude
	destination := points.MapPoints[1].Latitude
	destination += ", "
	destination += points.MapPoints[1].Longtitude
	cont, _ := maps.NewClient(maps.WithAPIKey(apiKey))
	r := &maps.DirectionsRequest{
		Origin:      origin,
		Destination: destination,
	}
	resp, _, err := cont.Directions(context.Background(), r)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"google api": &resp[0]})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": &resp[0]})
}