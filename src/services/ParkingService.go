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
	MaxParkingTime = 5 // seconds
)

type ParkingService struct {
	ParkingSpots   chan int
	EntryExitMutex sync.Mutex
	WG             sync.WaitGroup
	OutputMutex    sync.Mutex
	WaitingVehicles int
}

// NewParkingService creates a new ParkingService with a set number of parking spots
func NewParkingService(totalSpaces int) *ParkingService {
	ps := &ParkingService{
		ParkingSpots: make(chan int, totalSpaces),
	}
	for i := 0; i < totalSpaces; i++ {
		ps.ParkingSpots <- i
	}
	return ps
}

func (ps *ParkingService) ParkVehicle(vehicle *models.VehicleStatus) {
	vehicle.Status = "Esperando espacio"
	vehicle.NotifyObservers()

	spot := <-ps.ParkingSpots

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

	ps.ParkingSpots <- spot
	vehicle.Status = "Ha salido"
	vehicle.NotifyObservers()
}
