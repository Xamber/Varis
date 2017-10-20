package varis

import (
	"fmt"
	"math"
	"os"
)

// activation function (sigmoid)
func activation_sigmoid(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

// Derivative function
func derivative_sigmoid(x float64) float64 {
	var fx = activation_sigmoid(x)
	return fx * (1 - fx)
}

// Function for sum all elements in slice
func sum(data []float64) float64 {
	var result float64
	for _, i := range data {
		result += i
	}
	return result
}

func debug(format string, a ...interface{}) {
	if DEBUG == true {
		fmt.Fprintf(os.Stdout, format, a...)
	}
}
