package varis

import (
	"crypto/rand"
	"fmt"
)

// Function for sum all elements in slice.
func sum(data []float64) float64 {
	var result float64
	for _, i := range data {
		result += i
	}
	return result
}

func generate_uuid() string {
	b := make([]byte, 16)
	rand.Read(b)
	uuid := fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}
