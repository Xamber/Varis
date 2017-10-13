package main

import (
	"math/rand"
)

type Neuron interface {
	AddOutputSynapse(syn *Synapse)
	AddInputSynapse(syn *Synapse)

	GetOutputSynapses() []*Synapse
	GetInputSynapses() []*Synapse

	Handle(value float64)
	Broadcast(value float64)

	CollectSignals() []float64
	Activation() float64
	Train(delta float64) float64
}

type LiveNeuron interface {
	Alive()
}

type RedirectNeuron interface {
	GetOutput() chan float64
}

type BaseNeuron struct {
	bias        float64
	input       float64
	cache       float64
	inSynapses  []*Synapse
	outSynapses []*Synapse
}

func (n *BaseNeuron) AddOutputSynapse(syn *Synapse) {
	n.outSynapses = append(n.outSynapses, syn)
}

func (n *BaseNeuron) AddInputSynapse(syn *Synapse) {
	n.inSynapses = append(n.inSynapses, syn)
}

func (n *BaseNeuron) GetOutputSynapses() []*Synapse {
	return n.outSynapses
}

func (n *BaseNeuron) GetInputSynapses() []*Synapse {
	return n.inSynapses
}

func (n *BaseNeuron) Handle(value float64) {
	n.Broadcast(value)
}

func (n *BaseNeuron) Broadcast(value float64) {
	for o := range n.outSynapses {
		n.outSynapses[o].in <- value
	}
}

func (n *BaseNeuron) CollectSignals() []float64 {
	inputSignals := make([]float64, len(n.inSynapses))
	for i := range inputSignals {
		inputSignals[i] = <-n.inSynapses[i].out
	}
	return inputSignals
}

func (n *BaseNeuron) Activation() float64 {
	inputSignals := n.CollectSignals()

	value := sum(inputSignals) + n.bias
	n.input = value

	outputSignal := activation_sigmoid(value)
	n.cache = value
	return outputSignal
}

func (n *BaseNeuron) Train(delta float64) float64 {
	neuronDelta := delta * derivative_sigmoid(n.input)
	n.bias += neuronDelta

	for _, s := range n.inSynapses {
		s.ChangeWeight(neuronDelta)
	}

	return neuronDelta

}

type InputNeuron struct {
	BaseNeuron
}

func CreateInputNeuron() *InputNeuron {
	neuron := InputNeuron{BaseNeuron{bias: rand.Float64()}}
	return &neuron
}

type HiddenNeuron struct {
	BaseNeuron
}

func CreateHiddenNeuron() *HiddenNeuron {
	neuron := HiddenNeuron{BaseNeuron{bias: rand.Float64()}}
	return &neuron
}

func (n *HiddenNeuron) Alive() {
	for {
		value := n.Activation()
		n.Broadcast(value)
	}
}

type OutputNeuron struct {
	BaseNeuron
	output chan float64
}

func CreateOutputNeuron() *OutputNeuron {
	neuron := OutputNeuron{BaseNeuron{bias: rand.Float64()}, make(chan float64)}
	return &neuron
}

func (n *OutputNeuron) GetOutput() chan float64 {
	return n.output
}

func (n *OutputNeuron) Alive() {
	for {
		value := n.Activation()
		n.output <- value
	}
}
