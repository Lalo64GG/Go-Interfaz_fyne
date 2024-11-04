package models

type Parking struct {
    Capacity      int
    Spots         chan struct{}
    EntryExitLock chan struct{}
}

func NewParking(capacity int) *Parking {
    return &Parking{
        Capacity:      capacity,
        Spots:         make(chan struct{}, capacity),
        EntryExitLock: make(chan struct{}, 1),
    }
}
