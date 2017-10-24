package varis

import (
	"math/rand"
)

type Neuron interface {
	getConnection() *connection
	getCache() float64

	train(delta float64)
	live()
}

type baseNeuron struct {
	conn connection

	bias  float64
	cache float64
}

func (n *baseNeuron) getConnection() *connection {
	return &n.conn
}

func (n *baseNeuron) getCache() float64 {
	return n.cache
}

func (n *baseNeuron) train(neuronDelta float64) {
	n.bias += neuronDelta
	for _, s := range n.getConnection().inSynapses {
		s.weight += s.cache * neuronDelta
	}
}

type inputNeuron struct {
	baseNeuron
}

func createInputNeuron() *inputNeuron {
	neuron := inputNeuron{baseNeuron{bias: rand.Float64()}}
	return &neuron
}

func (n *inputNeuron) live() {
	// Empty
}

type hiddenNeuron struct {
	baseNeuron
}

func createHiddenNeuron() *hiddenNeuron {
	neuron := hiddenNeuron{baseNeuron{bias: rand.Float64()}}
	return &neuron
}

func (n *hiddenNeuron) live() {
	var signals []float64
	for {
		signals = n.getConnection().collectSignals()
		n.cache = sum(signals) + n.bias
		output := ACTIVATION_FUNCTION(n.cache)
		n.getConnection().broadcastSignals(output)
	}
}

type outputNeuron struct {
	baseNeuron
	output chan float64
}

func createOutputNeuron(outputChan chan float64) *outputNeuron {
	neuron := outputNeuron{baseNeuron{bias: rand.Float64()}, outputChan}
	return &neuron
}

func (n *outputNeuron) live() {
	var signals []float64
	for {
		signals = n.getConnection().collectSignals()
		n.cache = sum(signals) + n.bias
		output := ACTIVATION_FUNCTION(n.cache)
		n.output <- output
	}
}
