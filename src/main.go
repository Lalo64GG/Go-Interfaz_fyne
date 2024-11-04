package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/lalo64/parking-simulator/src/controllers"
	"github.com/lalo64/parking-simulator/src/services"
	"github.com/lalo64/parking-simulator/src/views"
)

const (
	totalSpaces   = 20
	totalVehicles = 100
)

func main() {
	myApp := app.New()

	// Inicializar servicios y controlador
	parkingService := services.NewParkingService(totalSpaces)
	vehicleService := services.NewVehicleService(totalVehicles)
	controller := controllers.NewVehicleController(parkingService, vehicleService)

	// Crear la vista del estacionamiento
	views.CreateParkingView(myApp, parkingService, vehicleService, controller)

	myApp.Run()
}
