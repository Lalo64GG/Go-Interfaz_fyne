package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/lalo64/parking-simulator/src/controllers"
	"github.com/lalo64/parking-simulator/src/services"
	"strconv"
	"fyne.io/fyne/v2"
)

const (
	totalSpaces   = 20
	totalVehicles = 50
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Simulador de Estacionamiento")

	outputLabel := widget.NewLabel("Iniciando simulador...")
	availableSpacesLabel := widget.NewLabel("Espacios disponibles: " + strconv.Itoa(totalSpaces))
	waitingVehiclesLabel := widget.NewLabel("Vehículos esperando: 0")

	parkingService := services.NewParkingService(totalSpaces)
	vehicleService := services.NewVehicleService(totalVehicles)
	controller := controllers.NewVehicleController(parkingService, vehicleService)

	// Configuración de la tabla con tamaños ajustados
	table := widget.NewTable(
		func() (int, int) { return len(vehicleService.Vehicles), 3 },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(i widget.TableCellID, o fyne.CanvasObject) {
			label := o.(*widget.Label)
			vehicle := vehicleService.GetVehicle(i.Row)
			switch i.Col {
			case 0:
				label.SetText("Vehículo " + strconv.Itoa(vehicle.ID))
			case 1:
				label.SetText(strconv.Itoa(vehicle.WaitTime) + " segundos")
			case 2:
				label.SetText(vehicle.Status)
			}
		},
	)

	// Ajusta el ancho de las columnas para mejorar la visualización
	table.SetColumnWidth(0, 250) // Columna para ID de vehículo
	table.SetColumnWidth(1, 200) // Columna para tiempo de espera
	table.SetColumnWidth(2, 300) // Columna para estado

	// Ajusta el tamaño mínimo del contenedor de la tabla
	scrollContainer := container.NewVScroll(table)
	scrollContainer.SetMinSize(fyne.NewSize(800, 500)) // Tamaño mínimo de la tabla

	startButton := widget.NewButton("Iniciar Simulación", func() {
		for i := 0; i < totalVehicles; i++ {
			go controller.StartVehicleSimulation(i)
		}
	})

	myWindow.SetContent(container.NewVBox(
		outputLabel,
		availableSpacesLabel,
		waitingVehiclesLabel,
		startButton,
		scrollContainer,
	))

	// Aumenta el tamaño de la ventana para mejor visibilidad
	myWindow.Resize(fyne.NewSize(1000, 800))
	myWindow.ShowAndRun()
}
