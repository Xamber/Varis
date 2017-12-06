package varis

import (
	"math"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var PrintCalculation = false
var PrintTrainLog = false

type neuronFunction func(x float64) float64

// ACTIVATION store default activation function.
var ACTIVATION neuronFunction = func(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

// DEACTIVATION store default deactivation function.
var DEACTIVATION neuronFunction = func(x float64) float64 {
	var fx = ACTIVATION(x)
	return fx * (1 - fx)
}
