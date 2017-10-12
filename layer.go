package main

import "fmt"

type Layerer interface {
	GetNeurons() []Neuron
	GetCountOfNeurons() int
	GetNeuronByIndex(index int) Neuron
	PrintInfo()
	RunAllNeurons()
}

type Layer struct {
	neurons []Neuron
}

func ConnectLayers(now Layerer, next Layerer) {
	for i := range now.GetNeurons() {
		for o := range next.GetNeurons() {
			CreateSynapse(now.GetNeuronByIndex(i), next.GetNeuronByIndex(o))
		}
	}
}

func (l *Layer) GetNeurons() []Neuron {
	return l.neurons
}

func (l *Layer) GetCountOfNeurons() int {
	return len(l.neurons)
}

func (l *Layer) GetNeuronByIndex(index int) Neuron {
	return l.neurons[index]
}

func (l *Layer) RunAllNeurons() {
	for _, neuron := range (l.GetNeurons()) {
		go neuron.(LiveNeuron).Alive()
	}
}

func (l *Layer) PrintInfo() {
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


func CreateInputLayer(neuronsCount int) *Layer {
	layer := Layer{}
	for i := 0; i < neuronsCount; i++ {
		layer.neurons = append(layer.neurons, Neuron(CreateInputNeuron()))
	}
	return &layer
}


func CreateHiddenLayer(neuronsCount int) *Layer {
	layer := Layer{}
	for i := 0; i <= neuronsCount; i++ {
		layer.neurons = append(layer.neurons, Neuron(CreateHiddenNeuron()))
	}
	return &layer
}


func CreateOutputLayer(neuronsCount int) *Layer {
	layer := Layer{}
	for i := 0; i < neuronsCount; i++ {
		layer.neurons = append(layer.neurons, Neuron(CreateOutputNeuron()))
	}
	return &layer
}
