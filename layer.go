package main

type Layer struct {
	neurons []Neuron
}

func (l *Layer) GetNeurons() *[]Neuron {
	return &l.neurons
}
