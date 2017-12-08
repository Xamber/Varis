package varis

import (
	"fmt"
	"math/rand"
)

// Perceptron implement Neural Network Perceptron by collect layers with Neurons and input/output channels.
type Perceptron struct {
	layers [][]Neuron
	input  []chan float64
	output []chan float64
}

// Calculate run network calculations by wait signals from input channels and send signals to output array of chan.
func (n *Perceptron) Calculate(input Vector) Vector {
	if len(input) != len(n.layers[0]) {
		panic("Check count of input value")
	}

	input.broadcast(n.input)
	output := collectVector(n.output)

	if PrintCalculation == true {
		fmt.Printf("Input: %v Output: %v\n", input, output)
	}

	return output
}

// RunNeurons create goroutines for all Neuron in Perceptron.
func (n *Perceptron) RunNeurons() {
	for _, l := range n.layers {
		for _, neuron := range l {
			go neuron.live()
		}
	}
}

// ConnectLayers create all to all connection between layers.
func (n *Perceptron) ConnectLayers() {
	for l := 0; l < len(n.layers)-1; l++ {
		now := n.layers[l]
		next := n.layers[l+1]
		for i := range now {
			for o := range next {
				ConnectNeurons(now[i], next[o], rand.Float64())
			}
		}
	}
}

// CreatePerceptron make new Perceptron NN with count of neurons in each Layer.
func CreatePerceptron(layers ...int) Perceptron {

	network := Perceptron{}
	network.input = make([]chan float64, 0)
	network.output = make([]chan float64, 0)

	for index, neurons := range layers {
		layer := []Neuron{}
		for i := 0; i < neurons; i++ {
			var neuron Neuron
			switch index {
			case 0:
				channel := make(chan float64)
				neuron = INeuron(rand.Float64(), channel)
				network.input = append(network.input, channel)
			case len(layers) - 1:
				channel := make(chan float64)
				neuron = ONeuron(rand.Float64(), channel)
				network.output = append(network.output, channel)
			default:
				neuron = HNeuron(rand.Float64())
			}

			layer = append(layer, neuron)
		}
		network.layers = append(network.layers, layer)
	}

	network.ConnectLayers()
	network.RunNeurons()

	return network
}
