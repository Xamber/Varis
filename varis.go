package varis

import (
	"math"
	"math/rand"
)

var ACTIVATION func(x float64) float64 = func(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

var DEACTIVATION func(x float64) float64 = func(x float64) float64 {
	var fx = ACTIVATION(x)
	return fx * (1 - fx)
}

func CreateNetwork(layers ...int) Network {
	network := Network{Output: make([]chan float64, 0)}
	for index, neurons := range layers {
		layer := &layer{}
		for i := 0; i < neurons; i++ {
			var neuron Neuroner
			switch index {
			case 0:
				neuron = CreateInputNeuron()
			case len(layers) - 1:
				neuron = CreateOutputNeuron(&network)
			default:
				neuron = CreateHiddenNeuron()
			}
			layer.AddNeuron(neuron)
		}
		network.AddLayer(layer)
	}

	network.ConnectLayers()
	network.RunNeurons()

	return network
}

func CreateInputNeuron() *Neuron {
	return &Neuron{bias: rand.Float64()}
}

func CreateHiddenNeuron() *Neuron {
	neuron := Neuron{bias: rand.Float64()}
	neuron.callbackFunc = neuron.conn.broadcastSignals
	return &neuron
}

func CreateOutputNeuron(network *Network) *Neuron {
	outputChan := make(chan float64)
	neuron := Neuron{bias: rand.Float64()}
	neuron.callbackFunc = func(value float64) {
		outputChan <- value
	}
	(*network).Output = append((*network).Output, outputChan)
	return &neuron
}
