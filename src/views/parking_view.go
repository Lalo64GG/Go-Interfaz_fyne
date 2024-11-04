// views/parking_view.go
package views

import (
	"image/color"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/lalo64/parking-simulator/src/controllers"
	"github.com/lalo64/parking-simulator/src/services"
)

const (
	totalSpaces   = 20
	totalVehicles = 100
)

func CreateParkingView(app fyne.App, parkingService *services.ParkingService, vehicleService *services.VehicleService, controller *controllers.VehicleController) {
	myWindow := app.NewWindow("Simulador de Estacionamiento")

	outputLabel := widget.NewLabel("Iniciando simulador...")
	availableSpacesLabel := widget.NewLabel("Espacios disponibles: " + strconv.Itoa(parkingService.AvailableSpaces()))
	waitingVehiclesLabel := widget.NewLabel("Vehículos esperando: " + strconv.Itoa(vehicleService.WaitingVehicles()))

	// Crear un grid para representar los espacios de estacionamiento
	parkingSlots := make([]*canvas.Rectangle, totalSpaces)
	gridContainer := container.NewGridWithColumns(5) // 5 espacios por fila

	for i := 0; i < totalSpaces; i++ {
		slotRectangle := canvas.NewRectangle(color.Gray{Y: 200}) // Color gris para espacios vacíos
		slotRectangle.SetMinSize(fyne.NewSize(80, 80))           // Tamaño de cada espacio
		parkingSlots[i] = slotRectangle
		gridContainer.Add(slotRectangle)
	}

	// Función para actualizar el color de los espacios según el estado del vehículo
	updateParkingSlots := func() {
		for i, slot := range parkingSlots {
			if parkingService.IsSpaceOccupied(i) {
				slot.FillColor = color.RGBA{R: 255, G: 0, B: 0, A: 255} // Rojo si está ocupado
			} else {
				slot.FillColor = color.Gray{Y: 200} // Gris si está libre
			}
			slot.Refresh() // Refrescar color para actualizar la UI
		}
		availableSpacesLabel.SetText("Espacios disponibles: " + strconv.Itoa(parkingService.AvailableSpaces()))
		waitingVehiclesLabel.SetText("Vehículos esperando: " + strconv.Itoa(vehicleService.WaitingVehicles()))
	}

	// Agregar observadores a cada vehículo para actualizar la UI cuando cambia su estado
	for _, vehicle := range vehicleService.Vehicles {
		vehicle.AddObserver(func() {
			updateParkingSlots()
		})
	}

	// Botón para iniciar la simulación
	startButton := widget.NewButton("Iniciar Simulación", func() {
		for i := 0; i < totalVehicles; i++ {
			go controller.StartVehicleSimulation(i)
			time.Sleep(500 * time.Millisecond) // Simular entrada gradual de vehículos
		}
	})

	myWindow.SetContent(container.NewVBox(
		outputLabel,
		availableSpacesLabel,
		waitingVehiclesLabel,
		startButton,
		gridContainer,
	))

	myWindow.Resize(fyne.NewSize(1000, 800))
	myWindow.Show()
}