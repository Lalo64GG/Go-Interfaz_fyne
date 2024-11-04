// utils/RandomTime.go
package utils

import "math/rand"

// RandomIntInRange returns a random integer within a specified range
func RandomIntInRange(min, max int) int {
	return rand.Intn(max-min+1) + min
}
