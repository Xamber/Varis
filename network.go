package varis

import (
	"math/rand"
)

// Network impliment Neural Network by collect layers with Neurons, output channel for store signals from output Layer.
type Network struct {
	layers [][]*Neuron
	output []chan float64
}

// AddLayer add Layer to Network.
func (n *Network) AddLayer(layer []*Neuron) {
	n.layers = append(n.layers, layer)
}

// Calculate run network calculations, and wait signals in output array of chan.
func (n *Network) Calculate(input []float64) []float64 {

	if len(input) != len(n.layers[0]) {
		panic("Check count of input value")
	}

	for i, n := range n.layers[0] {
		n.conn.broadcastSignals(input[i])
	}

	output := make([]float64, len(n.output))

	for i := range output {
		output[i] = <-n.output[i]
	}

	return output
}

// RunNeurons create goroutine for all Neuron in Network.
func (n *Network) RunNeurons() {
	for _, l := range n.layers {
		for _, neuron := range l {
			go neuron.live()
		}
	}
}

// ConnectLayers create all to all connection between layers.
func (n *Network) ConnectLayers() {
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
