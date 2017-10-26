package varis

import "math"

// Function for sum all elements in slice.
func sum(data []float64) float64 {
	var result float64
	for _, i := range data {
		result += i
	}
	return result
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}
