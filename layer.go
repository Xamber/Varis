package varis

import "fmt"

type Layerer interface {
	GetNeurons() []Neuron
	GetCountOfNeurons() int
	GetNeuronByIndex(index int) Neuron
	PrintInfo()
}

type Layer struct {
	neurons []Neuron
}

func CreateLayer() *Layer {
	return &Layer{}
}

func ConnectLayers(now Layerer, next Layerer) {
	for i := range now.GetNeurons() {
		for o := range next.GetNeurons() {
			CreateSynapse(now.GetNeuronByIndex(i), next.GetNeuronByIndex(o))
		}
	}
}

func (l *Layer) AddNeuron(neuron Neuron) {
	l.neurons = append(l.neurons, neuron)
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
