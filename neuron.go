package main

import (
	"math/rand"
	"fmt"
)

type Neuron interface {
	AddOutputSynapse(syn *Synapse)
	AddInputSynapse(syn *Synapse)

	GetOutputSynapses() []*Synapse
	GetInputSynapses() []*Synapse

	Handle(value float64)
	Broadcast(value float64)

	CollectSignals() []float64
}

type LiveNeuron interface {
	Alive()
}

type BaseNeuron struct {
	bias        float64
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
	fmt.Println("Hidden Neuron Run")
}

type OutputNeuron struct {
	BaseNeuron
}

func CreateOutputNeuron() *OutputNeuron {
	neuron := OutputNeuron{BaseNeuron{bias: rand.Float64()}}
	return &neuron
}

func (n *OutputNeuron) Alive() {
	fmt.Println("Output Neuron Run")
}
