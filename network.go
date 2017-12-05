package varis

import (
	"math/rand"
)

// Network impliment Neural Network by collect layers with Neurons, output channel for store signals from output Layer.
type Network struct {
	layers [][]*Neuron
	input  []chan float64
	output []chan float64
}

// Calculate run network calculations, and wait signals in output array of chan.
func (n *Network) Calculate(input Vector) Vector {
	if len(input) != len(n.layers[0]) {
		panic("Check count of input value")
	}

	input.Broadcast(n.input)
	output := CollectVector(n.output)

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

// CreateNetwork make new NN with count of neurons in each Layer.
func CreateNetwork(layers ...int) Network {

	network := Network{}
	network.input = make([]chan float64, 0)
	network.output = make([]chan float64, 0)

	for index, neurons := range layers {
		layer := []*Neuron{}
		for i := 0; i < neurons; i++ {

			var neuron *Neuron
			var channel chan float64

			switch index {
			case 0:
				neuron, channel = CreateNeuron(InputNeuron, rand.Float64())
				network.input = append(network.input, channel)
			case len(layers) - 1:
				neuron, channel = CreateNeuron(OutputNeuron, rand.Float64())
				network.output = append(network.output, channel)
			default:
				neuron, _ = CreateNeuron(HiddenNeuron, rand.Float64())
			}

			layer = append(layer, neuron)
		}
		network.layers = append(network.layers, layer)
	}

	network.ConnectLayers()
	network.RunNeurons()

	return network
}
