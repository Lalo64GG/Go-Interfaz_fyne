package controllers

import (
    "github.com/lalo64/parking-simulator/src/models"
)

func NewParkingController(capacity int) *models.Parking {
    return models.NewParking(capacity)
}
