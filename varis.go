package varis

import (
	"math"
	"math/rand"
	"time"
)

// Type for activation and deactivation functions. It should receive float64 and return float64
type neuronFunction func(x float64) float64

// PrintCalculation logs all calculate calls (print input and output).
var PrintCalculation = false

// Seed rand package
func init() {
	rand.Seed(time.Now().UnixNano())
}

// ACTIVATION store default activation function.
var ACTIVATION neuronFunction = func(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

// DEACTIVATION store default deactivation function.
var DEACTIVATION neuronFunction = func(x float64) float64 {
	var fx = ACTIVATION(x)
	return fx * (1 - fx)
}
