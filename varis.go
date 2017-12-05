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

// CreateNetwork make new NN with count of neurons in each Layer.
func CreateNetwork(layers ...int) Network {

	network := Network{}
	network.Output = make([]chan float64, 0)

	for index, neurons := range layers {
		layer := []*Neuron{}
		for i := 0; i < neurons; i++ {
			// Standart neuron implimentation
			// callFunc  is neuron.connection.broadcast
			// We will overwrite
			var neuron = &Neuron{weight: rand.Float64()}
			neuron.callbackFunc = neuron.conn.broadcastSignals

			switch index {
			case 0:
				// Input layer
				// Standart neuron implimentation without callFunc
				neuron.callbackFunc = nil
			case len(layers) - 1:
				// Output layer
				// Need to create output channel to redirect Neuron output to NetworkOutput
				outputChan := make(chan float64)
				neuron.callbackFunc = func(value float64) {
					outputChan <- value
				}
				network.Output = append(network.Output, outputChan)
			}
			layer = append(layer, neuron)
		}
		network.AddLayer(layer)
	}

	network.ConnectLayers()
	network.RunNeurons()

	return network
}
