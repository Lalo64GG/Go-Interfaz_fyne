// services/ParkingService.go
package services

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
	"github.com/lalo64/parking-simulator/src/models"
)

const (
	MinParkingTime = 3 // seconds
	MaxParkingTime = 8 // seconds
)

type ParkingService struct {
	ParkingSpots   chan int // Canal que representa los espacios de estacionamiento disponibles
	EntryExitMutex sync.Mutex
	WG             sync.WaitGroup
	OutputMutex    sync.Mutex
	WaitingVehicles int               // Número de vehículos en espera
	occupiedSpots   map[int]bool      // Mapa para rastrear qué espacios están ocupados
	mu              sync.Mutex        // Mutex para proteger acceso a occupiedSpots
}

func NewParkingService(totalSpaces int) *ParkingService {
	ps := &ParkingService{
		ParkingSpots: make(chan int, totalSpaces),
		occupiedSpots: make(map[int]bool),
	}
	for i := 0; i < totalSpaces; i++ {
		ps.ParkingSpots <- i
		ps.occupiedSpots[i] = false // Inicializar todos los espacios como desocupados
	}
	return ps
}

func (ps *ParkingService) ParkVehicle(vehicle *models.VehicleStatus) {
	ps.WaitingVehicles++
	vehicle.Status = "Esperando espacio"
	vehicle.NotifyObservers()

	spot := <-ps.ParkingSpots // Espera hasta que un espacio esté disponible
	ps.WaitingVehicles--

	ps.mu.Lock()
	ps.occupiedSpots[spot] = true // Marcar el espacio como ocupado
	ps.mu.Unlock()

	vehicle.Status = "Entrando"
	vehicle.NotifyObservers()
	time.Sleep(200 * time.Millisecond) // Simulate entry time

	ps.EntryExitMutex.Lock()
	time.Sleep(200 * time.Millisecond) // Simulate gate access time
	ps.EntryExitMutex.Unlock()

	parkingTime := rand.Intn(MaxParkingTime-MinParkingTime+1) + MinParkingTime
	vehicle.Status = fmt.Sprintf("Estacionado por %d segundos", parkingTime)
	vehicle.WaitTime = parkingTime
	vehicle.NotifyObservers()
	time.Sleep(time.Duration(parkingTime) * time.Second)

	ps.EntryExitMutex.Lock()
	time.Sleep(200 * time.Millisecond) // Simulate exit time
	ps.EntryExitMutex.Unlock()

	ps.mu.Lock()
	ps.occupiedSpots[spot] = false // Marcar el espacio como libre
	ps.mu.Unlock()
	ps.ParkingSpots <- spot // Devolver el espacio al estacionamiento
	vehicle.Status = "Ha salido"
	vehicle.NotifyObservers()
}

// Método para verificar si un espacio está ocupado
func (ps *ParkingService) IsSpaceOccupied(spot int) bool {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	return ps.occupiedSpots[spot]
}

// Método para obtener la cantidad de espacios disponibles
func (ps *ParkingService) AvailableSpaces() int {
	return len(ps.ParkingSpots)
}
