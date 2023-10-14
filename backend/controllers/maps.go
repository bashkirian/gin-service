package controllers

import (
	"net/http"
	_ "strconv"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/bashkirian/gin-service/models"
	"googlemaps.github.io/maps"
	"golang.org/x/net/context"
)

const apiKey = "AIzaSyDZ-FFbuQ0xhk1ArrZW8zZ8LdpUuIDsD0g"

// GET /map/route
// Get route from two points and distance 
func FindRoute(c *gin.Context) error {
	var points models.MapPointPayload
	if err := c.ShouldBindJSON(&points); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return fmt.Errorf("bind map json: %w", err)
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
	if err == nil {
		return fmt.Errorf("directions: %w", err)
	}
	c.JSON(http.StatusOK, gin.H{"data": &resp[0]})
	return nil
}