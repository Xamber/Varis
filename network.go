package varis

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Networker interface {
	Calculate(input ...float64) []float64
	AddLayer()
}

type Network struct {
	Layers []Layerer
	Output []chan float64
}

func CreateNetwork(layers ...int) Network {

	network := Network{Output: make([]chan float64, 0)}

	inputLayerIndex := 0
	outputLayerIndex := len(layers) - 1

	for index, neurons := range layers {

		layer := &layer{}

		for i := 0; i < neurons; i++ {
			neuron := Neuron{bias: rand.Float64()}
			neuron.activationFunc = activation_sigmoid

			switch index {
			case inputLayerIndex:
				// Empty for inputLayer
			case outputLayerIndex:
				outputChan := make(chan float64)
				neuron = Neuron{bias: rand.Float64()}
				neuron.activationFunc = activation_sigmoid
				neuron.callbackFunc = func(path chan float64) func(value float64) {
					return func(value float64) {
						path <- value
					}
				}(outputChan)
				network.Output = append(network.Output, outputChan)
			default:
				neuron.activationFunc = activation_sigmoid
				neuron.callbackFunc = neuron.conn.broadcastSignals
			}

			layer.AddNeuron(&neuron)

		}
		network.AddLayer(layer)
	}

	for l := 0; l < len(network.Layers)-1; l++ {
		now := network.Layers[l]
		next := network.Layers[l+1]
		for i := range now.getNeurons() {
			for o := range next.getNeurons() {
				createSynapse(now.getNeuronByIndex(i), next.getNeuronByIndex(o))
			}
		}
	}

	for _, l := range network.Layers {
		for _, neuron := range l.getNeurons() {
			go neuron.live()
		}
	}

	return network
}

func (n *Network) AddLayer(layer Layerer) {
	n.Layers = append(n.Layers, Layerer(layer))
}

func (n *Network) getInputLayer() Layerer {
	return n.Layers[0]
}

func (n *Network) getOutputLayer() Layerer {
	return n.Layers[len(n.Layers)-1]
}

func (n *Network) Calculate(input ...float64) []float64 {

	if len(input) != n.getInputLayer().getCountOfNeurons() {
		panic("Check count of input value")
	}

	for i, n := range n.getInputLayer().getNeurons() {
		n.getConnection().broadcastSignals(input[i])
	}

	output := make([]float64, len(n.Output))

	for i := range output {
		output[i] = <-n.Output[i]
	}

	return output
}
