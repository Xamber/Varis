package varis

import (
	"math/rand"
)

// Network impliment Neural Network by collect Layers with Neurons, output channel for store signals from output Layer.
type Network struct {
	Layers [][]*Neuron
	Output []chan float64
}

// AddLayer add Layer to Network.
func (n *Network) AddLayer(layer []*Neuron) {
	n.Layers = append(n.Layers, layer)
}

// Calculate run network calculations, and wait signals in Output array of chan.
func (n *Network) Calculate(input []float64) []float64 {

	if len(input) != len(n.Layers[0]) {
		panic("Check count of input value")
	}

	for i, n := range n.Layers[0] {
		n.conn.broadcastSignals(input[i])
	}

	output := make([]float64, len(n.Output))

	for i := range output {
		output[i] = <-n.Output[i]
	}

	return output
}

// RunNeurons create goroutine for all Neuron in Network.
func (n *Network) RunNeurons() {
	for _, l := range n.Layers {
		for _, neuron := range l {
			go neuron.live()
		}
	}
}

// ConnectLayers create all to all connection between layers.
func (n *Network) ConnectLayers() {
	for l := 0; l < len(n.Layers)-1; l++ {
		now := n.Layers[l]
		next := n.Layers[l+1]
		for i := range now {
			for o := range next {
				ConnectNeurons(now[i], next[o], rand.Float64())
			}
		}
	}
}
