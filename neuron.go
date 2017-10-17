package main

import (
	"math/rand"
)

type HaveInput interface {
	AddInputSynapse(syn *Synapse)
	GetInputSynapses() []*Synapse
}

type HaveOutput interface {
	AddOutputSynapse(syn *Synapse)
	GetOutputSynapses() []*Synapse

	Handle(value float64)
	Broadcast(value float64)
}

type Trainable interface {
	Train(delta float64)
}

type Workable interface {
	CollectSignals() []float64
	Activation() float64
	Deactivation() float64
}

type LiveNeuron interface {
	Alive()
}

type Redirectable interface {
	GetOutput() chan float64
}

type Neuron interface {
	HaveInput
	HaveOutput
	Workable
	Trainable
}

type BaseNeuron struct {
	bias        float64
	cache       float64
	inSynapses  []*Synapse
	outSynapses []*Synapse
}

type InputNeuron struct {
	BaseNeuron
}

type HiddenNeuron struct {
	BaseNeuron
}

type OutputNeuron struct {
	BaseNeuron
	output chan float64
}

func CreateBaseNeuron() BaseNeuron {
	return BaseNeuron{bias: rand.Float64()}
}

func CreateInputNeuron() *InputNeuron {
	neuron := InputNeuron{CreateBaseNeuron()}
	return &neuron
}

func CreateHiddenNeuron() *HiddenNeuron {
	neuron := HiddenNeuron{CreateBaseNeuron()}
	return &neuron
}

func CreateOutputNeuron() *OutputNeuron {
	neuron := OutputNeuron{CreateBaseNeuron(), make(chan float64)}
	return &neuron
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
	n.cache = sum(n.CollectSignals()) + n.bias
	outputSignal := activation_sigmoid(n.cache)
	return outputSignal
}

func (n *BaseNeuron) Deactivation() float64 {
	return derivative_sigmoid(n.cache)
}

func (n *BaseNeuron) Train(neuronDelta float64) {
	n.bias += neuronDelta
	for _, s := range n.inSynapses {
		s.weight += s.cache * neuronDelta
	}
}

func (n *HiddenNeuron) Alive() {
	for {
		n.Broadcast(n.Activation())
	}
}

func (n *OutputNeuron) GetOutput() chan float64 {
	return n.output
}

func (n *OutputNeuron) Alive() {
	for {
		n.output <- n.Activation()
	}
}
