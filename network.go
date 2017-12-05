package varis

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Network impliment Neural Network by collect Layers with Neurons, output channel for store signals from output Layer.
type Network struct {
	Layers []Layer
	Output []chan float64
}

// AddLayer add Layer to Network.
func (n *Network) AddLayer(layer Layer) {
	n.Layers = append(n.Layers, layer)
}

// Calculate run network calculations, and wait signals in Output array of chan.
func (n *Network) Calculate(input []float64) []float64 {

	if len(input) != len(n.GetInputLayer()) {
		panic("Check count of input value")
	}

	for i, n := range n.GetInputLayer() {
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
				ConnectNeurons(now[i], next[o], generate_uuid(), rand.Float64())
			}
		}
	}
}

// GetInputLayer from Network.
func (n *Network) GetInputLayer() Layer {
	return n.Layers[0]
}

// GetOutputLayer from Network.
func (n *Network) GetOutputLayer() Layer {
	return n.Layers[len(n.Layers)-1]
}

// CreateInputNeuron make new Neuron without callback function.
func (n *Network) createInputNeuron(uuid string, bias float64) *Neuron {
	return &Neuron{bias: bias, uuid: uuid}
}

// CreateHiddenNeuron make new Neuron with default callback function.
func (n *Network) createHiddenNeuron(uuid string, bias float64) *Neuron {
	neuron := Neuron{bias: bias, uuid: uuid}
	neuron.callbackFunc = neuron.conn.broadcastSignals
	return &neuron
}

// CreateOutputNeuron make new Neuron with redirect output and append new channel to network.Output.
func (n *Network) createOutputNeuron(uuid string, bias float64) *Neuron {
	outputChan := make(chan float64)
	neuron := Neuron{bias: bias, uuid: uuid}
	neuron.callbackFunc = func(value float64) {
		outputChan <- value
	}
	n.Output = append(n.Output, outputChan)
	return &neuron
}
