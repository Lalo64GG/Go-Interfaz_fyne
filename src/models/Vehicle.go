// models/VehicleStatus.go
package models

type VehicleStatus struct {
	ID        int
	Status    string
	WaitTime  int
	Observers []func() // Observers to update the GUI
}

// AddObserver adds a new observer to the VehicleStatus
func (vs *VehicleStatus) AddObserver(observer func()) {
	vs.Observers = append(vs.Observers, observer)
}

// NotifyObservers updates all observers with the new state
func (vs *VehicleStatus) NotifyObservers() {
	for _, observer := range vs.Observers {
		observer()
	}
}
