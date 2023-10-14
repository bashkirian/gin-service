package models

import "github.com/google/uuid"

type Client struct {
	ID          uuid.UUID
	LocationLat string `json:"location_lat"`
	LocationLon string `json:"location_lon"`
}
