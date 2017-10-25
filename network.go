package varis

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Network impliment Neural Network by collect layers with Neurons, output channel for store signals from output Layer.
type Network struct {
	Layers []Layerer
	Output []chan float64
}

// AddLayer add layer to Network.
func (n *Network) AddLayer(layer Layerer) {
	n.Layers = append(n.Layers, Layerer(layer))
}

// Calculate run network calculations, and wait signals in Output array of chan.
func (n *Network) Calculate(input ...float64) []float64 {

	if len(input) != n.GetInputLayer().getCountOfNeurons() {
		panic("Check count of input value")
	}

	for i, n := range n.GetInputLayer().getNeurons() {
		n.getConnection().broadcastSignals(input[i])
	}

	output := make([]float64, len(n.Output))

	for i := range output {
		output[i] = <-n.Output[i]
	}

	return output
}

// RunNeurons create goroutine for all neuron in Network.
func (n *Network) RunNeurons() {
	for _, l := range n.Layers {
		for _, neuron := range l.getNeurons() {
			go neuron.live()
		}
	}
}

// ConnectLayers create all to all connection between layers.
func (n *Network) ConnectLayers() {
	for l := 0; l < len(n.Layers)-1; l++ {
		now := n.Layers[l]
		next := n.Layers[l+1]
		for i := range now.getNeurons() {
			for o := range next.getNeurons() {
				createSynapse(now.getNeuronByIndex(i), next.getNeuronByIndex(o))
			}
		}
	}
}

// GetInputLayer from Network.
func (n *Network) GetInputLayer() Layerer {
	return n.Layers[0]
}

// GetOutputLayer from Network.
func (n *Network) GetOutputLayer() Layerer {
	return n.Layers[len(n.Layers)-1]
}
