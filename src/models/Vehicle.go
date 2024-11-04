package models

type VehicleStatus struct {
	ID        int
	Status    string
	WaitTime  int
	Observers []func() // Observers to update the GUI
}

// AddObserver adds a new observer to the VehicleStatus
func (vs *VehicleStatus) AddObserver(observer func()) { //  agrega una función observadora que se ejecutará cuando cambie el estado del vehículo.
	vs.Observers = append(vs.Observers, observer)
}

// NotifyObservers updates all observers with the new state
func (vs *VehicleStatus) NotifyObservers() { // ejecuta todas las funciones en Observers, lo cual es útil para actualizar la interfaz de usuario cuando el estado cambia.
	for _, observer := range vs.Observers {
		observer()
	}
}
