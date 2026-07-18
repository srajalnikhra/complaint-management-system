package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// GenerateOTP creates a 6-digit random code used to check email ownership for password resets.
func GenerateOTP() string {
	// Seed the pseudo-random generator with the current time nano offset.
	rand.Seed(time.Now().UnixNano())

	// Produce a six digit random number code.
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}
