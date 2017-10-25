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
		layer := &layer{}
		for i := 0; i < neurons; i++ {
			var neuron Neuroner
			switch index {
			case 0: // Input layer
				neuron = createInputNeuron()
			case len(layers) - 1: // Output layer
				neuron = createOutputNeuron(&network)
			default: // Hidden layer
				neuron = createHiddenNeuron()
			}
			layer.AddNeuron(neuron)
		}
		network.AddLayer(layer)
	}

	network.ConnectLayers()
	network.RunNeurons()

	return network
}

// CreateInputNeuron make new neuron without callback function.
func createInputNeuron() *neuron {
	return &neuron{bias: rand.Float64()}
}

// CreateHiddenNeuron make new neuron with default callback function.
func createHiddenNeuron() *neuron {
	neuron := neuron{bias: rand.Float64()}
	neuron.callbackFunc = neuron.conn.broadcastSignals
	return &neuron
}

// CreateOutputNeuron make new neuron with redirect output and append new channel to network.Output.
func createOutputNeuron(network *Network) *neuron {
	outputChan := make(chan float64)
	neuron := neuron{bias: rand.Float64()}
	neuron.callbackFunc = func(value float64) {
		outputChan <- value
	}
	(*network).Output = append((*network).Output, outputChan)
	return &neuron
}
