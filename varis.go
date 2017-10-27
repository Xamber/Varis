package varis

import (
	"math"
	"math/rand"
)

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

// CreateNetwork make new NN with count of neurons in each layer.
func CreateNetwork(layers ...int) Network {
	network := Network{Output: make([]chan float64, 0)}
	for index, neurons := range layers {
		layer := layer{}
		for i := 0; i < neurons; i++ {
			var neuron Neuroner
			switch index {
			case 0: // Input layer
				neuron = network.createInputNeuron(generate_uuid(), rand.Float64())
			case len(layers) - 1: // Output layer
				neuron = network.createOutputNeuron(generate_uuid(), rand.Float64())
			default: // Hidden layer
				neuron = network.createHiddenNeuron(generate_uuid(), rand.Float64())
			}
			layer.AddNeuron(neuron)
		}
		network.AddLayer(layer)
	}

	network.ConnectLayers()
	network.RunNeurons()

	return network
}
