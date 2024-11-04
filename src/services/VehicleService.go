package services

import (
	"github.com/lalo64/parking-simulator/src/models"
)

type VehicleService struct {
	Vehicles []*models.VehicleStatus // Estados de todos los vehículos
}

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

// Método para contar la cantidad de vehículos en espera
func (vs *VehicleService) WaitingVehicles() int {
	count := 0
	for _, vehicle := range vs.Vehicles {
		if vehicle.Status == "Esperando espacio" {
			count++
		}
	}
	return count
}
