package main

import "fmt"

type Layer interface {
	GetNeurons() []Neuron
	PrintInfo()
}

type BaseLayer struct {
	neurons []Neuron
}

func (l *BaseLayer) GetNeurons() []Neuron {
	return l.neurons
}

func (l *BaseLayer) PrintInfo() {
	fmt.Println("Layer:")
	for index, neuron := range l.GetNeurons() {
		fmt.Println("    Neuron (", index, "): ", neuron)

		for _, synapse := range neuron.GetInputSynapses() {
			fmt.Println("        InSynapse: ", synapse)
		}

		for _, synapse := range neuron.GetOutputSynapses() {
			fmt.Println("        OutSynapse: ", synapse)
		}
	}
}

type InputLayer struct {
	BaseLayer
}

func CreateInputLayer(neuronsCount int) *InputLayer {
	layer := InputLayer{}

	for i := 0; i < neuronsCount; i++ {
		layer.neurons = append(layer.neurons, Neuron(CreateInputNeuron()))
	}

	return &layer
}

type HiddenLayer struct {
	BaseLayer
}

func CreateHiddenLayer(neuronsCount int) *HiddenLayer {
	layer := HiddenLayer{}

	for i := 0; i < neuronsCount; i++ {
		layer.neurons = append(layer.neurons, Neuron(CreateHiddenNeuron()))
	}

	return &layer
}

type OutputLayer struct {
	BaseLayer
}

func CreateOutputLayer(neuronsCount int) *OutputLayer {
	layer := OutputLayer{}

	for i := 0; i < neuronsCount; i++ {
		layer.neurons = append(layer.neurons, Neuron(CreateOutputNeuron()))
	}

	return &layer
}
