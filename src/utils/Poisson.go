package utils

import (
    "math/rand"
    "time"
)

func PoissonRandomTime() time.Duration {
    lambda := 300 // promedio de llegada en milisegundos
    return time.Duration(rand.ExpFloat64() * float64(lambda)) * time.Millisecond
}
