package main

import (
	"math/rand"
)

type Neuron interface {
	AddInputSynapse(syn *Synapse)
	AddOutputSynapse(syn *Synapse)

	GetInputSynapses() []*Synapse
	GetOutputSynapses() []*Synapse

	Handle(value float64)
	Broadcast(value float64)
	CollectSignals() []float64

	Activation() float64
	Deactivation() float64

	Train(delta float64)

	Alive()
}

type Redirectable interface {
	GetOutput() chan float64
}

type CoreNeuron struct {
	bias        float64
	cache       float64
	inSynapses  []*Synapse
	outSynapses []*Synapse
}

type InputNeuron struct {
	CoreNeuron
}

type HiddenNeuron struct {
	CoreNeuron
}

type OutputNeuron struct {
	CoreNeuron
	output chan float64
}

func CreateCoreNeuron() CoreNeuron {
	return CoreNeuron{bias: rand.Float64()}
}

func CreateInputNeuron() *InputNeuron {
	neuron := InputNeuron{CreateCoreNeuron()}
	return &neuron
}

func CreateHiddenNeuron() *HiddenNeuron {
	neuron := HiddenNeuron{CreateCoreNeuron()}
	return &neuron
}

func CreateOutputNeuron() *OutputNeuron {
	neuron := OutputNeuron{CreateCoreNeuron(), make(chan float64)}
	return &neuron
}

func (n *CoreNeuron) AddOutputSynapse(syn *Synapse) {
	n.outSynapses = append(n.outSynapses, syn)
}

func (n *CoreNeuron) AddInputSynapse(syn *Synapse) {
	n.inSynapses = append(n.inSynapses, syn)
}

func (n *CoreNeuron) GetOutputSynapses() []*Synapse {
	return n.outSynapses
}

func (n *CoreNeuron) GetInputSynapses() []*Synapse {
	return n.inSynapses
}

func (n *CoreNeuron) Handle(value float64) {
	n.Broadcast(value)
}

func (n *CoreNeuron) Broadcast(value float64) {
	for o := range n.outSynapses {
		n.outSynapses[o].in <- value
	}
}

func (n *CoreNeuron) CollectSignals() []float64 {
	inputSignals := make([]float64, len(n.inSynapses))
	for i := range inputSignals {
		inputSignals[i] = <-n.inSynapses[i].out
	}
	return inputSignals
}

func (n *CoreNeuron) Activation() float64 {
	n.cache = sum(n.CollectSignals()) + n.bias
	outputSignal := activation_sigmoid(n.cache)
	return outputSignal
}

func (n *CoreNeuron) Deactivation() float64 {
	return derivative_sigmoid(n.cache)
}

func (n *CoreNeuron) Train(neuronDelta float64) {
	n.bias += neuronDelta
	for _, s := range n.inSynapses {
		s.weight += s.cache * neuronDelta
	}
}

func (n *CoreNeuron) Alive() {
	// Empty
}

func (n *HiddenNeuron) Alive() {
	for {
		n.Broadcast(n.Activation())
	}
}

func (n *OutputNeuron) Alive() {
	for {
		n.output <- n.Activation()
	}
}

func (n *OutputNeuron) GetOutput() chan float64 {
	return n.output
}
