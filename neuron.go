package varis

import (
	"math/rand"
)

type Neuron interface {
	addInputSynapse(syn *synapse)
	addOutputSynapse(syn *synapse)

	getInputSynapses() []*synapse
	getOutputSynapses() []*synapse

	handle(value float64)
	broadcast(value float64)
	collectSignals() []float64

	activation() float64
	deactivation() float64

	train(delta float64)

	alive()
}

type Redirectable interface {
	GetOutput() chan float64
}

type coreNeuron struct {
	bias        float64
	cache       float64
	inSynapses  []*synapse
	outSynapses []*synapse
}

type inputNeuron struct {
	coreNeuron
}

type hiddenNeuron struct {
	coreNeuron
}

type outputNeuron struct {
	coreNeuron
	output chan float64
}

func createCoreNeuron() coreNeuron {
	return coreNeuron{bias: rand.Float64()}
}

func createInputNeuron() *inputNeuron {
	neuron := inputNeuron{createCoreNeuron()}
	return &neuron
}

func createHiddenNeuron() *hiddenNeuron {
	neuron := hiddenNeuron{createCoreNeuron()}
	return &neuron
}

func createOutputNeuron(outputChan chan float64) *outputNeuron {
	neuron := outputNeuron{createCoreNeuron(), outputChan}
	return &neuron
}

func (n *coreNeuron) addOutputSynapse(syn *synapse) {
	n.outSynapses = append(n.outSynapses, syn)
}

func (n *coreNeuron) addInputSynapse(syn *synapse) {
	n.inSynapses = append(n.inSynapses, syn)
}

func (n *coreNeuron) getOutputSynapses() []*synapse {
	return n.outSynapses
}

func (n *coreNeuron) getInputSynapses() []*synapse {
	return n.inSynapses
}

func (n *coreNeuron) handle(value float64) {
	n.broadcast(value)
}

func (n *coreNeuron) broadcast(value float64) {
	for o := range n.outSynapses {
		n.outSynapses[o].in <- value
	}
}

func (n *coreNeuron) collectSignals() []float64 {
	inputSignals := make([]float64, len(n.inSynapses))
	for i := range inputSignals {
		inputSignals[i] = <-n.inSynapses[i].out
	}
	return inputSignals
}

func (n *coreNeuron) activation() float64 {
	n.cache = sum(n.collectSignals()) + n.bias
	outputSignal := activation_sigmoid(n.cache)
	return outputSignal
}

func (n *coreNeuron) deactivation() float64 {
	return derivative_sigmoid(n.cache)
}

func (n *coreNeuron) train(neuronDelta float64) {
	n.bias += neuronDelta
	for _, s := range n.inSynapses {
		s.weight += s.cache * neuronDelta
	}
}

func (n *coreNeuron) alive() {
	// Empty
}

func (n *hiddenNeuron) alive() {
	for {
		n.broadcast(n.activation())
	}
}

func (n *outputNeuron) alive() {
	for {
		n.output <- n.activation()
	}
}
