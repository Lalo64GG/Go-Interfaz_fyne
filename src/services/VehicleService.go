// services/VehicleService.go
package services

import (
	"github.com/lalo64/parking-simulator/src/models"
)

type VehicleService struct {
	Vehicles []*models.VehicleStatus
}

// NewVehicleService initializes a new VehicleService with a set number of vehicles
func NewVehicleService(totalVehicles int) *VehicleService {
	vs := &VehicleService{}
	for i := 0; i < totalVehicles; i++ {
		vs.Vehicles = append(vs.Vehicles, &models.VehicleStatus{ID: i, Status: "Esperando", WaitTime: 0})
	}
	return vs
}

func (vs *VehicleService) GetVehicle(id int) *models.VehicleStatus {
	return vs.Vehicles[id]
}
