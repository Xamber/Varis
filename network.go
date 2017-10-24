package varis

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var DEBUG bool = false

type Network struct {
	Layers []Layerer
	Output []chan float64
}

func CreateNetwork(layers ...int) Network {

	network := Network{Output: make([]chan float64, 0)}

	inputLayerIndex := 0
	outputLayerIndex := len(layers) - 1

	for index, neurons := range layers {

		layer := CreateLayer()

		for i := 0; i < neurons; i++ {
			var neuron Neuron

			switch index {
			case inputLayerIndex:
				neuron = createInputNeuron()
			case outputLayerIndex:
				outputChan := make(chan float64)
				neuron = createOutputNeuron(outputChan)
				network.Output = append(network.Output, outputChan)
			default:
				neuron = createHiddenNeuron()
			}

			layer.addNeuron(neuron)

		}
		network.addLayer(layer)
	}

	for l := 0; l < len(network.Layers)-1; l++ {
		now := network.Layers[l]
		next := network.Layers[l+1]
		ConnectLayers(now, next)
	}

	network.runAllNeuron()

	return network
}

func (n *Network) addLayer(layer Layerer) {
	n.Layers = append(n.Layers, Layerer(layer))
}

func (n *Network) getInputLayer() Layerer {
	return n.Layers[0]
}

func (n *Network) getOutputLayer() Layerer {
	return n.Layers[len(n.Layers)-1]
}

func (n *Network) runAllNeuron() {
	for _, l := range n.Layers {
		for _, neuron := range l.getNeurons() {
			go neuron.alive()
		}
	}
}

func (n *Network) Calculate(input ...float64) []float64 {

	if len(input) != n.getInputLayer().getCountOfNeurons() {
		panic("Check count of input value")
	}

	for i, n := range n.getInputLayer().getNeurons() {
		n.broadcast(input[i])
	}

	output := make([]float64, len(n.Output))

	for i := range output {
		output[i] = <-n.Output[i]
	}

	return output
}
