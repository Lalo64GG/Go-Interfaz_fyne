// controllers/VehicleController.go
package controllers

import (

	"github.com/lalo64/parking-simulator/src/services"
)

type VehicleController struct {
	ParkingService *services.ParkingService
	VehicleService *services.VehicleService
}

func NewVehicleController(parkingService *services.ParkingService, vehicleService *services.VehicleService) *VehicleController {
	return &VehicleController{
		ParkingService: parkingService,
		VehicleService: vehicleService,
	}
}

func (vc *VehicleController) StartVehicleSimulation(id int) { // Obtiene un vehículo por su ID y llama a ParkVehicle en ParkingService para simular el proceso de estacionamiento.

	vehicle := vc.VehicleService.GetVehicle(id)
	vc.ParkingService.ParkVehicle(vehicle)
}
