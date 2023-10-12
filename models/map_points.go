package models

type MapPoint struct {
	Latitude string `json:"latitude" binding:"required"`
	Longtitude string `json:"longtitude" binding:"required"`
}

type MapPointPayload struct {
	MapPoints []MapPoint `binding:"dive"`
}